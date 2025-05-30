package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
	"time"
	"unicode"

	"github.com/emicklei/melrose/core"

	"github.com/emicklei/melrose/dsl"
)

var tmplSource = `---
title: "Language"
description: "All language functions grouped by creation,compostion and audio control."
lead: "All language functions grouped by creation,compostion and audio control."
date: 2021-03-05T12:21:01+02:00
lastmod: 2025-05-09T12:21:01+02:00
draft: false
images: []
menu: 
  docs:
    parent: "reference"
weight: 20
toc: true
---

## Structure

### Expressions

Musical objects are created, composed and played using the <strong>melrōse</strong> tool by evaluating expressions.
Expressions use any of the predefined functions (creation,composition,audio control).
By assigning an expression to a variable name, you can use that expression by its name to compose other objects.

### variables

Variable names must start with a non-digit character and can have zero or more characters in [a-z A-Z _ 0-9].
An assignment "=" is used to create a variable.
To delete a variable, assign it to the special value "nil".

### comment

Use "//" to add comment, either on a new line or and the end of an expression.

### programs

Programs are composed of comment lines, assigments and expressions.
Each must be separated by a newline or a semicolon ;.
Large expressions may span multiple lines only when lines after the first are indented by a Tab or 4 Space characters.

## Creation functions
{{range .Core}}
- <a href="#{{.Anchor}}">{{.Title}}</a>{{end}}

## Composition functions
{{range .Composer}}
- <a href="#{{.Anchor}}">{{.Title}}</a>{{end}}

## Audio control functions
{{range .Audio}}
- <a href="#{{.Anchor}}">{{.Title}}</a>{{end}}

{{range .All}}
### {{.Title}}
<a name="{{.Anchor}}"></a>
{{.Description}}

> {{.Syntax}}
	
{{ backticks }}javascript{{ range .Examples }}
{{ . }}
{{ end }}{{ backticks }}
{{end}}


##### generated by dsl-md.go on {{.Created.Format "Jan 02, 2006"}}
`

type DocumentedFunction struct {
	Title            string
	ShortDescription string
	Syntax           string
	Description      string
	Examples         []string
	Anchor           string
	Alias            string
}

type GroupedFunctions struct {
	Created  time.Time
	Core     []DocumentedFunction
	Composer []DocumentedFunction
	Audio    []DocumentedFunction
	All      []DocumentedFunction
}

func main() {
	gf := GroupedFunctions{Created: time.Now()}
	ctx := core.PlayContext{
		VariableStorage: dsl.NewVariableStore(),
		LoopControl:     core.NoLooper,
	}
	for k, each := range dsl.EvalFunctions(ctx) {
		// draft functions do not have a Title
		if len(each.Title) == 0 {
			continue
		}
		// TODO needed?
		// aliasses do not have a separate doc section
		if unicode.IsUpper(rune(k[0])) {
			continue
		}
		// skip entries without description
		if len(each.Description) == 0 {
			continue
		}
		// skip alias entries
		if each.Alias == k {
			continue
		}
		title := k
		if len(each.Alias) > 0 && each.Alias != title {
			title = fmt.Sprintf("%s (%s)", k, each.Alias)
		}
		df := DocumentedFunction{
			Title:            title,
			ShortDescription: each.Title,
			Description:      firstUpcaseAndDot(each.Description),
			Examples:         strings.Split(each.Samples, "\n"),
			Anchor:           k,
			Syntax:           each.HumanizedTemplate(),
			Alias:            each.Alias,
		}
		if each.ControlsAudio {
			gf.Audio = append(gf.Audio, df)
		} else {
			if each.IsCore {
				gf.Core = append(gf.Core, df)
			} else {
				gf.Composer = append(gf.Composer, df)
			}
		}
		gf.All = append(gf.All, df)
	}
	sort.Slice(gf.Core, func(i, j int) bool { return gf.Core[i].Title < gf.Core[j].Title })
	sort.Slice(gf.Composer, func(i, j int) bool { return gf.Composer[i].Title < gf.Composer[j].Title })
	sort.Slice(gf.Audio, func(i, j int) bool { return gf.Audio[i].Title < gf.Audio[j].Title })
	sort.Slice(gf.All, func(i, j int) bool { return gf.All[i].Title < gf.All[j].Title })

	name := "../../content/docs/reference/dsl.md"
	out, err := os.Create(name)
	checkErr(err)
	defer out.Close()
	tmpl := template.New("dsl")
	tmpl.Funcs(template.FuncMap{
		"backticks": func() string {
			return "```"
		},
	})
	tmpl.Parse(tmplSource)
	err = tmpl.Execute(out, gf)
	checkErr(err)
	where, _ := filepath.Abs(name)
	fmt.Println("generated:", where)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func firstUpcaseAndDot(s string) string {
	b := bytes.Buffer{}
	last := ' '
	for i, each := range []rune(s) {
		if i == 0 {
			b.WriteRune(unicode.ToUpper(each))
		} else {
			b.WriteRune(each)
		}
		last = each
	}
	if last != '.' {
		b.WriteRune('.')
	}
	return b.String()
}
