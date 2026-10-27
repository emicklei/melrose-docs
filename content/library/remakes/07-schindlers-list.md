---
title: "07-schindlers-list"
menu:
  library:
    parent: remakes
weight: 70
---

<ul>
<li>Schindler&#39;s List - Hans Zimmer</li>
<li>Ableton - Orchestral Sweep Pad</li>
<li>Garageband - Luminous Sweep Pad, Mellow Poly</li>
</ul>

<div class="mel-source">

```javascript
// Schindler's List - Hans Zimmer
// Ableton - Orchestral Sweep Pad
// Garageband - Luminous Sweep Pad, Mellow Poly

bpm(50)

s1 = sequence('a3 8a 8d 8a 8d 8b_ 8a 8f 8a3 8f 8g3 8f 8g 2a')
s2 = fraction(8,sequence('a d a d c5 b_ a g c b_3 e g b_ a g f'))
s3 = sequence('8g 8b_3 8g 8b_3 8a 8g 8f 8g 2.a')
s4 = sequence('a3 8g 8b_3 8g 8b_3 8a 8g 8f 8e 2d 2e')
s5 = sequence('8f 8b_3 8f 8b_3 8a 8g 8f 8e 2.d =')

loop(s1,s2,s3,s4,s5)
```

</div>

<div class="copy-source"><button class="btn-clipboard btn-copy-source btn btn-sm btn-link" data-clipboard-target=".docs-content .highlight code" title="Copy to clipboard" aria-label="Copy to clipboard"><svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-copy" aria-hidden="true"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg></button></div>
