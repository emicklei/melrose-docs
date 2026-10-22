---
title: "03-karmino"
menu:
  library:
    parent: melodies
weight: 30
---

<ul>
<li>Karmino by Nil Zork</li>
<li>Logic Pro X, Luminous Tines</li>
<li>Analog Lab, Bells</li>
</ul>

<div class="mel-source">

```javascript
// Karmino by Nil Zork
// Logic Pro X, Luminous Tines
// Analog Lab, Bells

it = iterator(6,1,4,-1) // delta semitones
pm = pitchmap('1:0,1:7,1:12,1:15,1:19,1:24',note('c')) // six notes
sm = fraction(8,resequence('1 4 2 4 3 6 5 4',pm)) // note sequence of eights
lp_sq = loop(pitch(it,sm),next(it)) // loop the sequence and change pitch on every iteration
```

</div>

<div class="copy-source"><button class="btn-clipboard btn-copy-source btn btn-sm btn-link" data-clipboard-target=".docs-content .highlight code" title="Copy to clipboard" aria-label="Copy to clipboard"><svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-copy" aria-hidden="true"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg></button></div>
