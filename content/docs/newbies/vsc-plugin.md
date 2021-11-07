---
title: "VSC Plugin"
description: "Microsoft Visual Studio Code Plugin."
lead: "Microsoft Visual Studio Code Plugin."
date: 2020-11-16T13:59:39+01:00
lastmod: 2020-11-16T13:59:39+01:00
draft: false
images: []
menu:
  docs:
    parent: "newbies"
weight: 120
toc: true
---

### Install

See [instructions →](https://marketplace.visualstudio.com/items?itemName=EMicklei.melrose-for-vscode) on how to install this plugin after you have installed Visual Studio Code.
This editor extension works with `.mel` and `.melrose` files.
The syntax of the program uses the [Melrōse Lanugage →](/docs/reference/dsl/).

### ⌘+e : Evaluate

To evaluate a single line statement or expression, the cursor must be on that line and then use `cmd+e`.
You can also evaluate source you have selected using the same shortcut `cmd+e`.
To evaluate a program, you need to select all the source and use `cmd+e`.

### ⌘+3 : Play

To play a single line statement or expression, use `cmd+3`.
You can also evaluate the function `play(...)`.

### ⌘+5 : Stop

To stop a running loop or listener, use `cmd+5`.
You can also evaluate the function `stop(...)`.

### ⌘+2 : Inspect

To inspect a variable or a function, just hover with your mouse pointer above its name.
To explicitly inspect the value of an expression, use `cmd+2`.

### ⌘+k : Stop all sounds

To stop sounds being played, including loops, use `cmd+k`.

## Editing tips

### comment

Lines that start with `//` are not evaluated ; these are commment lines.

	// this is comment

Lines can have inline comment at the end.

	s = note('C#') // C sharp, Octave 4

### multiline

A statement can span multiple lines, each line after the first must be indented by either a TAB or 4 spaces.

	  y = sequence('F#2 
	  [TAB]C#3 F#3 A3 C# F#')
	  x = sequence('A 
	  [SPACE][SPACE][SPACE][SPACE]A5 A6')