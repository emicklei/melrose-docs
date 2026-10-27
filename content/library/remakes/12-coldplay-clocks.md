---
title: "12-coldplay-clocks"
menu:
  library:
    parent: remakes
weight: 120
---

<ul>
<li>Coldplay - Clocks</li>
<li>GarageBand: High Rise Pluck</li>
</ul>

<div class="mel-source">

```javascript
// Coldplay - Clocks
// GarageBand: High Rise Pluck

bpm(131)

// 3+3+2 arpeggiated groupings
bar_eb  = sequence('8eb5 8bb4 8g4  8eb5 8bb4 8g4  8eb5 8bb4')
bar_bbm = sequence('8db5 8bb4 8f4  8db5 8bb4 8f4  8db5 8bb4')
bar_fm  = sequence('8c5  8ab4 8f4  8c5  8ab4 8f4  8c5  8ab4')

// The full 4-bar riff (Bbm plays twice)
clocks_riff = join(bar_eb, bar_bbm, bar_bbm, bar_fm)

// Bass roots
bass_riff = sequence('1eb2 1bb1 1bb1 1f2')

l1 = loop(channel(1, clocks_riff))
l2 = loop(channel(2, bass_riff))

sync(l1, l2)
```

</div>

<div class="copy-source"><button class="btn-clipboard btn-copy-source btn btn-sm btn-link" data-clipboard-target=".docs-content .highlight code" title="Copy to clipboard" aria-label="Copy to clipboard"><svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-copy" aria-hidden="true"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg></button></div>
