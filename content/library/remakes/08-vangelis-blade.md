---
title: "08-vangelis-blade"
menu:
  library:
    parent: remakes
weight: 80
---

<ul>
<li>Vangelis Blade</li>
<li>Analog Lab, Deckards&#39;s Dx</li>
<li>GarageBand, Mellow Poly Synth</li>
</ul>

<div class="mel-source">

```javascript
// Vangelis Blade
// Analog Lab, Deckards's Dx
// GarageBand, Mellow Poly Synth

bpm(50)

s1 = sequence('8c++ 16d .e+ 16f 8g 16a g')
s2 = sequence('8a++ 16b .c5+ 16b 8a 16g a')
s3 = sequence('8g++ 16f .e+ 16d 8e 16g d')
s4 = sequence('8c++ 16b3 .a3+ 16g3 8b3 16c d')
r = note('8=')
l1 = join(s1,r,s2,r,s3,r,s4,r)

loop(l1)
```

</div>

<div class="copy-source"><button class="btn-clipboard btn-copy-source btn btn-sm btn-link" data-clipboard-target=".docs-content .highlight code" title="Copy to clipboard" aria-label="Copy to clipboard"><svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-copy" aria-hidden="true"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg></button></div>
