---
title: "MCP Server"
description: "Tool to compose music using conversation in AI"
lead: "Tool to compose music using conversation in AI."
date: 2025-05-09T19:53:58+01:00
lastmod: 2025-05-09T19:53:58+01:00
draft: false
weight: 50 
contributors: []
---

`melrose-mcp` is a tool that uses the [MCP](https://modelcontextprotocol.io/) protocol to receive `melrōse` expressions to play.
You install this tool locally on your machine and connect it from an MCP client such as Claude Desktop, Goose, VSC etc.

### example

Prompt:

```
play a tune from Claude Debussy
```

AI: 

```
I'll create a small segment inspired by Claude Debussy's impressionist style using the Melrose language.
```

MCP server:

    {
        `expression`: `sequence('(F4 A4 D5) 4(F4 A4 D5) (G4 B4 E5) 4(G4 B4 E5) (A4 C5 F5) 4(A4 C5 F5) (G4 B4 E5) 4(G4 B4 E5) (F4 A4 D5) 8= 8(E4 G4 C5) 4(F4 A4 D5)')`
    }

See [melrose-mcp](https://github.com/emicklei/melrose-mcp) for details how to install and use it.
