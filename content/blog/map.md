---
title: "Introducing 'map'"
description: "Use map to transform a collection of sequences using a replacement function"
lead: "Use map to transform a collection of sequences (requires v0.55+)."
date: 2024-09-10T09:58:42+02:00
lastmod: 2024-09-10T09:58:42+02:00
draft: false
weight: 50 
contributors: []
---

The function `map` is used to create a composition by applying a transformation on each sequence in a collection.

{{< highlight javascript "linenos=table,linenostart=1" >}}
a1 = sequence('G_3 D_ G_ A D_5 G_5')
a2 = sequence('D_3 A_3 D_ E A_ D_5')
a3 = sequence('E3 B3 E A_ B E5 ')
a4 = sequence('B2 G_3 B3 E_ G_ B')

c = map( join(a1,a3,a4,a2), fraction(8,resequence('1 4 2 4 3 6 5 4', _ )) )
loop(c)
{{< / highlight >}}

### line 1..4

Four sequences are created that play a nice combination of arpeggios.

### line 6

To create a collection, all sequences are composed in a `join`.
The transformation is done using a `fraction` on top of a `resequence`.
Note that the `resequence` uses the temporary variable called `_`.
This composition will the produce the following flat sequence:

    sequence('8G_3 8A 8D_ 8A 8G_ 8G_5 8D_5 8A 8E3 8A_ 8B3 8A_ 8E 8E5 8B 8A_ 8B2 8E_ 8G_3 8E_ 8B3 8B 8G_ 8E_ 8D_3 8E 8A_3 8E 8D_ 8D_5 8A_ 8E') 