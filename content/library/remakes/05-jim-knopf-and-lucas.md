---
title: "05-jim-knopf-and-lucas"
menu:
  library:
    parent: remakes
weight: 50
---

<ul>
<li>Jim Knopf and Lucas the Engine Driver</li>
<li>Music by Michael Ende</li>
<li>Translated by Ernest Micklei using Melrōse</li>
<li>Garageband: Square Bells</li>
</ul>

<div class="mel-source">

```javascript
// Jim Knopf and Lucas the Engine Driver
// Music by Michael Ende
// Translated by Ernest Micklei using Melrōse
// Garageband: Square Bells

bpm(120)
s1 = sequence('.c5 8b_ a_ e_ c-- e_-- a_-- =')
s2 = sequence('.d_5 8c5 b_ f d_-- f-- b_-- =')
s3 = sequence('.b_ 8a_ g e_ b_3-- e_-- g-- =')
line = stretch(0.5,join(s1,s1,s2,s2,s3,s3,s1,s1))
loop(line)
```

</div>

<div class="copy-source"><button class="btn-clipboard btn-copy-source btn btn-sm btn-link" data-clipboard-target=".docs-content .highlight code" title="Copy to clipboard" aria-label="Copy to clipboard"><svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-copy" aria-hidden="true"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg></button></div>
