---
title: "11-airforce"
menu:
  library:
    parent: melodies
weight: 110
---

<ul>
<li>Airforce by Nil Zork</li>
<li>Analog Lab Pro: Blue Bells</li>
<li>Garageband: Dreamy Bells</li>
</ul>

<div class="mel-source">

```javascript
// Airforce by Nil Zork
// Analog Lab Pro: Blue Bells
// Garageband: Dreamy Bells

bpm(140)

p1 = iterator(0,-2,1)
b0 = fraction(8,sequence('f# c#5 f#5 d a d5 a3 e a c# g# c#5'))  
b1 = pitch(p1,b0)
l1 = loop(b1,next(p1))
```

</div>

<div class="copy-source"><button class="btn-clipboard btn-copy-source btn btn-sm btn-link" data-clipboard-target=".docs-content .highlight code" title="Copy to clipboard" aria-label="Copy to clipboard"><svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-copy" aria-hidden="true"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg></button></div>
