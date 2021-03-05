---
title: "Iterating over resequence patterns"
description: ""
lead: ""
date: 2021-03-05T17:28:49+01:00
lastmod: 2021-03-05T17:28:49+01:00
draft: false
weight: 50 
contributors: []
---

{{< highlight javascript "linenos=table,linenostart=1" >}}

bpm(90)
s = sequence('c e g b')
i = iterator(
    '(1 2) (3 4)',
    '(1 3) (2 4)', 
    '1 (2 3) 4', 
    '(1 4) (2 3)')
m = resequence(i,s)
l = loop(m,next(i))

{{< / highlight >}}

Using `resequence` you can compose a new sequence with index-based entries.
The first compostion played by the loop is `(C E) (G B)`.