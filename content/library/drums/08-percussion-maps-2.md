---
title: "08-percussion-maps-2"
menu:
  library:
    parent: drums
weight: 80
---

<ul>
<li>Finale SongWriter 2012 Percussion Maps Reference</li>
<li><a href="https://usermanuals.finalemusic.com/SongWriter2012Win/Content/PercussionMaps.htm">https://usermanuals.finalemusic.com/SongWriter2012Win/Content/PercussionMaps.htm</a></li>
</ul>

<div class="mel-source">

```javascript
// Finale SongWriter 2012 Percussion Maps Reference
// https://usermanuals.finalemusic.com/SongWriter2012Win/Content/PercussionMaps.htm

bpm(85)

open = midi(16,46,60)
close = midi(16,42,60)
clap = midi(16,39,60)
kick = midi(16,36,80)

d1 = notemap('!..!..!..!..!..!',open)
d2 = notemap('..!..!..!..!..!.',kick)
d3 = notemap('.!..!..!..!..!..',close)
dset = channel(10,merge(d1,d2,d3))
lp_set = loop(dset)


// Big Room
kick = note('16f2++')
k1 = notemap('!!..!!..!.!.!...',kick)
lp_k1 = loop(k1)
```

</div>

<div class="copy-source"><button class="btn-clipboard btn-copy-source btn btn-sm btn-link" data-clipboard-target=".docs-content .highlight code" title="Copy to clipboard" aria-label="Copy to clipboard"><svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-copy" aria-hidden="true"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg></button></div>
