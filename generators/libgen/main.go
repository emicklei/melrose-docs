package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
)

var libraryDir = flag.String("l", "../../library", "directory with library sections of .mel files")

var docDir = flag.String("d", "../../content/library", "directory of generated documentation")

var templateFile = flag.String("t", "library-template.md", "template for each generated markdown file")

var sections = []string{"drums", "melodies"}

type page struct {
	Title    string
	Section  string
	Weight   int
	Comments []string
	Source   string
}

func main() {
	flag.Parse()
	tmpl, err := template.New(filepath.Base(*templateFile)).
		Funcs(template.FuncMap{"hasPrefix": strings.HasPrefix}).
		ParseFiles(*templateFile)
	if err != nil {
		log.Fatal(err)
	}
	for _, section := range sections {
		where, _ := filepath.Abs(filepath.Join(*libraryDir, section, "*.mel"))
		fmt.Println("scanning", where)
		matches, err := filepath.Glob(where)
		if err != nil {
			log.Fatal(err)
		}
		sources := map[string]bool{}
		for i, each := range matches {
			sources[strings.TrimSuffix(filepath.Base(each), ".mel")] = true
			generate(tmpl, section, each, (i+1)*10)
		}
		removeOrphans(section, sources)
	}
}

// removeOrphans deletes generated markdown files whose .mel source no longer exists.
// Files starting with "_" (such as _index.md) are maintained by hand and kept.
func removeOrphans(section string, sources map[string]bool) {
	generated, err := filepath.Glob(filepath.Join(*docDir, section, "*.md"))
	if err != nil {
		log.Fatal(err)
	}
	for _, each := range generated {
		name := filepath.Base(each)
		if strings.HasPrefix(name, "_") || sources[strings.TrimSuffix(name, ".md")] {
			continue
		}
		fmt.Println("removing", each)
		if err := os.Remove(each); err != nil {
			log.Fatal(err)
		}
	}
}

func generate(tmpl *template.Template, section, fileName string, defaultWeight int) {
	source, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	title := strings.TrimSuffix(filepath.Base(fileName), ".mel")
	p := page{
		Title:    title,
		Section:  section,
		Weight:   weightOf(title, defaultWeight),
		Comments: leadingComments(string(source)),
		Source:   strings.TrimRight(string(source), " \t\r\n"),
	}
	outputFile := filepath.Join(*docDir, section, title+".md")
	output, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()
	fmt.Println("generating", outputFile)
	if err := tmpl.Execute(output, p); err != nil {
		log.Fatal(err)
	}
}

// leadingComments returns the non-empty leading "//" comment lines of the source.
// The "//" (and surrounding whitespace) is stripped from each comment line.
func leadingComments(source string) []string {
	var comments []string
	for _, line := range strings.Split(strings.ReplaceAll(source, "\r\n", "\n"), "\n") {
		trimmed := strings.TrimSpace(line)
		if !strings.HasPrefix(trimmed, "//") {
			break
		}
		if text := strings.TrimSpace(strings.TrimPrefix(trimmed, "//")); text != "" {
			comments = append(comments, text)
		}
	}
	return comments
}

// weightOf uses a numeric filename prefix such as "01-" to order pages in the menu.
func weightOf(title string, defaultWeight int) int {
	prefix, _, found := strings.Cut(title, "-")
	if !found {
		return defaultWeight
	}
	n, err := strconv.Atoi(prefix)
	if err != nil {
		return defaultWeight
	}
	return n * 10
}
