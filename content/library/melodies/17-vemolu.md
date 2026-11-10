---
title: "17-vemolu"
menu:
  library:
    parent: melodies
weight: 170
---

<ul>
<li>Vemolu by Nil Zork</li>
<li>Garage: Muted Bass, FM Piano</li>
<li>Analog Lab Pro: Cool Down The Pad</li>
</ul>

<div class="mel-source">

```javascript
// Vemolu by Nil Zork
// Garage: Muted Bass, FM Piano
// Analog Lab Pro: Cool Down The Pad

a1 = sequence('G_3 D_ G_ A D_5 G_5')
a2 = sequence('D_3 A_3 D_ E A_ D_5')
a3 = sequence('E3 B3 E A_ B E5 ')
a4 = sequence('B2 G_3 B3 E_ G_ B') 
b1 = fraction(8,a1,a2,a3,a4)

it = iterator(
    '1 4 2 4 3 6 5 4',
    '1 (4 5 6) 2 (4 5 6) 3 (4 5 6) 2 (4 5 6)',
    '(1 6) (2 5) (3 4) (2 5) (3 4) (2 5)', 
    '(1 2 3) (4 5 6) (4 5 6) (4 5 6)',
    '(1 3 5) (2 4 6)',
    '(1 3) (2 4) (3 5) (4 6)',
    '(1 5) (2 6) (3 1) (4 2)'
    )
i1 = resequence(it,a1) // change the order of notes in a1 according to the indices of the iterator's current value
i2 = replace(i1,a1,a2) // replace in i1 all occurrences of a1 with a2
i3 = replace(i2,a1,a3)
i4 = replace(i3,a1,a4)
b6 = fraction(8,i1,i2,i3,i4) // set the note length to 1/8 for all sequences

loop(repeat(2,b6),note('2='),print(it),next(it)) // loop repeating each pattern twice with half a rest in between


// first 4 notes only per sequence
r1 = resequence('1 2 3 4',a1)
f1 = fraction(8,r1,replace(r1,a1,a2),replace(r1,a1,a3),replace(r1,a1,a4))

r2 = resequence('(1 5) (2 6) (3 1) (4 2)',a1)
f2 = replace(f1,r1,r2)

r3 = repeat(2,resequence('(1 3 5) (2 4 6)',a1))
f3 = replace(f1,r1,r3)

r4 = resequence('(1 2 3 4) (3 4 5 6) (2 3 4 5) (5 6 1 2)',a1)
f4 = replace(f1,r1,r4)

r5 = resequence('(1 2) (2 3) (3 4) (5 6)',a1)
f5 = replace(f1,r1,r5)

l1 = loop(f1,f1,f2,f2,f3,f3,f4,f4)


loop(f2)


// the best

a1 = sequence('G_3 D_ G_ A D_5 G_5')
a2 = sequence('D_3 A_3 D_ E A_ D_5')
a3 = sequence('E3 B3 E A_ B E5 ')
a4 = sequence('B2 G_3 B3 E_ G_ B')

c = map( join(a1,a2,a3,a4), fraction(8,resequence('1 4 2 4 3 6 5 4', _ )) )
loop(c)

d = map( join(a1,a2,a3,a4), fraction(8,resequence('(1 5) (2 6) (3 1) (4 2)', _ )) )
loop(d)
```

</div>

<div class="copy-source"><button class="btn-clipboard btn-copy-source btn btn-sm btn-link" data-clipboard-target=".docs-content .highlight code" title="Copy to clipboard" aria-label="Copy to clipboard"><svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-copy" aria-hidden="true"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg></button></div>

Copy the script and try it using <a href="https://play.melrōse.org" target="_blank">Melrōse Playground</a>.
