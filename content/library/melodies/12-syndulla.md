---
title: "12-syndulla"
menu:
  library:
    parent: melodies
weight: 120
---

<ul>
<li>Syndulla by Nil Zork</li>
<li>Analog V: Falling Pad</li>
<li>Garageband: Warm Arp Pad</li>
</ul>

<div class="mel-source">

```javascript
// Syndulla by Nil Zork
// Analog V: Falling Pad
// Garageband: Warm Arp Pad

bpm(120)

s1 = sequence('8c5 8e_5 8g5 8e_5 8f5 8b_ 8d5 8f5')
s2 =  sequence('8g 8b_ 8d5 8b_ 8e_5 8a_ 8c5 8e_5')
s3 = join(s1,s2)
s4 = transpose(-7,s3)

loop(s3,s4, note('='))
```

</div>

<div class="copy-source"><button class="btn-clipboard btn-copy-source btn btn-sm btn-link" data-clipboard-target=".docs-content .highlight code" title="Copy to clipboard" aria-label="Copy to clipboard"><svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-copy" aria-hidden="true"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg></button></div>
