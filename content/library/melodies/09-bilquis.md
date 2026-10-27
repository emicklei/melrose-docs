---
title: "09-bilquis"
menu:
  library:
    parent: melodies
weight: 90
---

<ul>
<li>Bilquis by Nil Zork</li>
<li>Analog Lab, Reverse Box</li>
<li>Wavestate: Dark Sonata Split</li>
</ul>

<div class="mel-source">

```javascript
// Bilquis by Nil Zork
// Analog Lab, Reverse Box
// Wavestate: Dark Sonata Split

bpm(120)
s = sequence(' e  g#  c#5  e5      f#  a  c#5 f#5     g#  b  d#5 g#5   a5 g#5 f#5 e5')
t = sequence('2e2    2c#2         2f#2   2c#3        2g#2   2d#3')
u = fraction(8,octave(0,merge(s,t)))
repeat(4,join(u,note('1=')))
loop(u)
```

</div>

<div class="copy-source"><button class="btn-clipboard btn-copy-source btn btn-sm btn-link" data-clipboard-target=".docs-content .highlight code" title="Copy to clipboard" aria-label="Copy to clipboard"><svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-copy" aria-hidden="true"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg></button></div>
