package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/djherbis/times"
)

var sampleDir = flag.String("m", "/Users/emicklei/Projects/github.com/emicklei/melrose-projects/public", "directory with sample scripts")

var docDir = flag.String("d", "/Users/emicklei/Projects/github.com/emicklei/melrose-docs/content/examples", "directory of generated documentation")

func main() {
	flag.Parse()
	matches, _ := filepath.Glob(filepath.Join(*sampleDir, "*.mel"))
	for _, each := range matches {
		generate(each)
	}
}

func generate(fileName string) {
	// take title from file ; meta may override
	title := filepath.Base(fileName)
	title = strings.Title(title[3 : len(title)-4])
	tags := ""

	info, err := times.Stat(fileName)
	if err != nil {
		log.Fatal(err)
	}

	input, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	output, err := os.Create(filepath.Join(*docDir, title+".md"))
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()

	scanner := bufio.NewScanner(input)
	readingMeta := false
	readingComment := false
	readingScript := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "//*" {
			readingMeta = !readingMeta
			// end of meta
			if !readingMeta {
				fmt.Fprintln(output, "---")
				fmt.Fprintf(output, "title: \"%s\"\n", title)
				fmt.Fprintf(output, "date: %s\n", info.BirthTime().Format("2006-01-02"))
				fmt.Fprintln(output, "draft: false")
				fmt.Fprintln(output, "---")
				if tags != "" {
					fmt.Fprintln(output, "\t", tags)
					fmt.Fprintln(output)
				}
				readingComment = true
			}
			continue
		}
		if readingMeta {
			if strings.HasPrefix(line, "//title:") {
				title = strings.TrimPrefix(line, "//title:")
			}
			if strings.HasPrefix(line, "//tags:") {
				tags = strings.TrimPrefix(line, "//tags:")
			}
			continue
		}
		if !readingScript {
			if line == "" {
				fmt.Fprintln(output)
				continue
			}
			if strings.HasPrefix(line, "//") {
				if readingComment {
					fmt.Fprintln(output, strings.TrimPrefix(line, "//"))
					continue
				}
			} else {
				// end of comment
				readingComment = false
				readingScript = true
				fmt.Fprintln(output, "```javascript")
			}
		}
		if readingScript {
			fmt.Fprintln(output, line)
		}
	}
	fmt.Fprintln(output, "```")

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
