---
title: "14-whitesail"
menu:
  library:
    parent: melodies
weight: 140
---

<ul>
<li>Whitesail by Nil Zork</li>
<li>Garageband: Mellow Poly</li>
</ul>

<div class="mel-source">

```javascript
// Whitesail by Nil Zork
// Garageband: Mellow Poly

bpm(120)

s = sequence('c d# g 8f 8d# f g# c5 8a# 8g# a# c#5 f5 8d#5 8c#5 c5 g# f =')
t = sequence('1c2           1f2            1a#1                 c2 g#2 f2 =')

ls = loop(dynamic('-',s))
lt = loop(dynamic('--',t))
st = sync(ls,lt)
```

</div>

<div class="copy-source"><button class="btn-clipboard btn-copy-source btn btn-sm btn-link" data-clipboard-target=".docs-content .highlight code" title="Copy to clipboard" aria-label="Copy to clipboard"><svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-copy" aria-hidden="true"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg></button></div>
