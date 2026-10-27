---
title: "06-myonak"
menu:
  library:
    parent: melodies
weight: 60
---

<ul>
<li>Myonak by Nil Zork</li>
<li>Logic: Tech House Vibes</li>
<li>Ableton: Airlite</li>
</ul>

<div class="mel-source">

```javascript
// Myonak by Nil Zork
// Logic: Tech House Vibes
// Ableton: Airlite

bpm(100)

s1=sequence('g d#5 d#5 d5 c5 b c5 d5')
s2=sequence('g f5 f5 d#5 d5 c5 d5 d#5')
s3=sequence('a# g5 g5 f5 d#5 d5 d#5 f5')
s4=sequence('a# f5 f5 d#5 d5 c5 d5 d#5')

f1 = fraction(8,join(s1,s2,s3,s4))
loop(f1)
```

</div>

<div class="copy-source"><button class="btn-clipboard btn-copy-source btn btn-sm btn-link" data-clipboard-target=".docs-content .highlight code" title="Copy to clipboard" aria-label="Copy to clipboard"><svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-copy" aria-hidden="true"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg></button></div>
