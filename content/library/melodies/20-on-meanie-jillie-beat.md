---
title: "20-on-meanie-jillie-beat"
menu:
  library:
    parent: melodies
weight: 200
---

<ul>
<li>On Meanie Jillie Beat by Nil Zork</li>
<li>Korg Wavestation: Meanie Jillie Beat</li>
</ul>

<div class="mel-source">

```javascript
// On Meanie Jillie Beat by Nil Zork
// Korg Wavestation: Meanie Jillie Beat

bpm(128)

g2 = chord('g/m/2')
d5 = chord('d5')
c5 = chord('c5/m')
g1 = chord('g/m/1')

rh = stretch(8,join(g2,d5,c5,g1))
lh = stretch(4,sequence('g3 g3 d3 d3 c3 c3 g3 g3'))

merge(lh,rh)
```

</div>

<div class="copy-source"><button class="btn-clipboard btn-copy-source btn btn-sm btn-link" data-clipboard-target=".docs-content .highlight code" title="Copy to clipboard" aria-label="Copy to clipboard"><svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-copy" aria-hidden="true"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg></button></div>

Copy the script and try it using <a href="https://play.melrōse.org" target="_blank">Melrōse Playground</a>.
