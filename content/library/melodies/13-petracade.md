---
title: "13-petracade"
menu:
  library:
    parent: melodies
weight: 130
---

<ul>
<li>Petracade by Nil Zork</li>
<li>Analog Lab, Majestico</li>
<li>Garageband: French Horns</li>
</ul>

<div class="mel-source">

```javascript
// Petracade by Nil Zork
// Analog Lab, Majestico
// Garageband: French Horns

bpm(120)
mel = fraction(1,chordsequence('c/m f3/m/2 b_3 e_3/2 a_3 d3/dim/2 g3 c3/m/2'))
lp_mel = loop(mel)

bass = dynamic('++',octave(-2,fraction(1,sequence('c5 f b_ e_ a_ d g c'))))
lp_bass = loop(bass)

// intro = chordsequence('c/m b_3 d3/o/1 g3')

s = sync(lp_mel,lp_bass)
```

</div>

<div class="copy-source"><button class="btn-clipboard btn-copy-source btn btn-sm btn-link" data-clipboard-target=".docs-content .highlight code" title="Copy to clipboard" aria-label="Copy to clipboard"><svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-copy" aria-hidden="true"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg></button></div>
