---
title: "06-four-on-the-floor"
menu:
  library:
    parent: drums
weight: 60
---

<ul>
<li>Garageband: SoCal</li>
<li><a href="https://docs.google.com/spreadsheets/d/19_3BxUMy3uy1Gb0V8Wc-TcG7q16Amfn6e8QVw4-HuD0/edit#gid=0">https://docs.google.com/spreadsheets/d/19_3BxUMy3uy1Gb0V8Wc-TcG7q16Amfn6e8QVw4-HuD0/edit#gid=0</a></li>
<li>Four on the Floor</li>
</ul>

<div class="mel-source">

```javascript
// Garageband: SoCal
// https://docs.google.com/spreadsheets/d/19_3BxUMy3uy1Gb0V8Wc-TcG7q16Amfn6e8QVw4-HuD0/edit#gid=0
// Four on the Floor

kick=note('c2') // 1
clap=note('d#2') // 2
openhi=note('a#2') // 3
closehi=note('g#2') // 4
snare=note('e2') // 5
R4=note('=') // 6
rim=note('f3') // 7
ride=midi(duration(kick),51,72) // 8

all = join(kick,clap,openhi,closehi,snare,R4,rim,ride) // 1..8

// Four on the floor
bpm(120)
drum1 =fraction(16,joinmap('1 6 3 6 (1 2 5) 6 3 6 1 4 3 6 (1 2 5) 6 3 6',all))
lp_drum1 = loop(channel(10,dynamic('--',drum1)))
```

</div>

<div class="copy-source"><button class="btn-clipboard btn-copy-source btn btn-sm btn-link" data-clipboard-target=".docs-content .highlight code" title="Copy to clipboard" aria-label="Copy to clipboard"><svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-copy" aria-hidden="true"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg></button></div>
