---
title: "05-synthetic-substitution-2"
menu:
  library:
    parent: drums
weight: 50
---

<ul>
<li>drum patterns</li>
<li><a href="https://docs.google.com/spreadsheets/d/19_3BxUMy3uy1Gb0V8Wc-TcG7q16Amfn6e8QVw4-HuD0/edit#gid=0">https://docs.google.com/spreadsheets/d/19_3BxUMy3uy1Gb0V8Wc-TcG7q16Amfn6e8QVw4-HuD0/edit#gid=0</a></li>
<li>Synthetic Substitution part 2</li>
</ul>

<div class="mel-source">

```javascript
// drum patterns
// https://docs.google.com/spreadsheets/d/19_3BxUMy3uy1Gb0V8Wc-TcG7q16Amfn6e8QVw4-HuD0/edit#gid=0
// Synthetic Substitution part 2

kick=note('c2')
snare=note('e2')
closehi=midi(4,42,72)
openhi=midi(4,46,72)
hibell=midi(4,53,72)
lobell=midi(4,52,72)
																
drum15 = merge(
    notemap('!.!....!.!!!...!',kick),
    notemap('....!.......!...',snare),
    notemap('!.!.!.!.!.!.!.!.',closehi),
    notemap('..!.............',openhi))
lp_drum15 = loop(channel(3,fraction(16,drum15)))
// export("synthetic-part2.mid",fraction(16,drum15))
```

</div>

<div class="copy-source"><button class="btn-clipboard btn-copy-source btn btn-sm btn-link" data-clipboard-target=".docs-content .highlight code" title="Copy to clipboard" aria-label="Copy to clipboard"><svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-copy" aria-hidden="true"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg></button></div>
