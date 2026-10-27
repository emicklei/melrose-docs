---
title: "09-einaudi"
menu:
  library:
    parent: remakes
weight: 90
---

<ul>
<li>Einaudi song</li>
<li>Analog Lab 4, American Grey Road</li>
<li>Analog Lab V, Enchanted Classic</li>
<li>Wavestate, Iceland Piano Pad</li>
</ul>

<div class="mel-source">

```javascript
// Einaudi song
// Analog Lab 4, American Grey Road
// Analog Lab V, Enchanted Classic
// Wavestate, Iceland Piano Pad

bpm(120)

p1 = fraction(8,sequence('f3 c f a c5 f5 a5'))
p1b = reverse(trim(1,1,p1))
p2 = fraction(8,sequence('e3 c e a c5 e5 a5'))
p2b = replace(p1b,p1,p2)
p3 = fraction(8,sequence('e_3 b_3 e_ g b_ e_5 g5'))
p3b = replace(p1b,p1,p3)
p4 = fraction(8,sequence('g3 b_3 d g b_ d5 g5'))
p4b = replace(p1b,p1,p4)

m = join(p1,p1b,p2,p2b,p3,p3b,p4,p4b)
// bpm(300);loop(m)

b = dynamic('--',sequence('1.f2 1.e2 1.e_2 1.g2'))
loop(sync(m,b))
```

</div>

<div class="copy-source"><button class="btn-clipboard btn-copy-source btn btn-sm btn-link" data-clipboard-target=".docs-content .highlight code" title="Copy to clipboard" aria-label="Copy to clipboard"><svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-copy" aria-hidden="true"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg></button></div>
