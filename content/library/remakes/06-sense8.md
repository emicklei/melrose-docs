---
title: "06-sense8"
menu:
  library:
    parent: remakes
weight: 60
---

<ul>
<li>Sense8 - Main Theme</li>
<li>Translated by Ernest Micklei using Melrōse</li>
<li>Garageband: Bright Synth Strings, 80s Wave Bells</li>
</ul>

<div class="mel-source">

```javascript
// Sense8 - Main Theme
// Translated by Ernest Micklei using Melrōse
// Garageband: Bright Synth Strings, 80s Wave Bells

bpm(120)
s1 = fraction(2,sequence('(g2 d3 g3 b3)'))
s2 = fraction(0.375,sequence('(a2 e2 a3 c#)'))
s3 = fraction(2,sequence('(a#2 f3 b_3 d)'))
s4 = fraction(2,sequence('(d#2 a#2 d#3 g3 b_3)'))
s5 = fraction(0.375,sequence('(f2 c3 f3 a3 c)'))
s6 = fraction(1,sequence('(f#2 c#3 f#3 b_3 c#)'))
theme = join(s1,s1,s2,s3,fraction(1,s3),s4,s4,s5,s6)
loop(theme)

// export('39-sense8',theme)
```

</div>

<div class="copy-source"><button class="btn-clipboard btn-copy-source btn btn-sm btn-link" data-clipboard-target=".docs-content .highlight code" title="Copy to clipboard" aria-label="Copy to clipboard"><svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-copy" aria-hidden="true"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg></button></div>
