---
title: "01-frere-jacques"
menu:
  library:
    parent: melodies
weight: 10
---

<ul>
<li>Uses `track` to play sequences together, starting at given bars.</li>
</ul>

<div class="mel-source">

```javascript
// Uses `track` to play sequences together, starting at given bars.

f1 = sequence('C D E C')
f2 = sequence('E F 2G')
f3 = sequence('8G 8A 8G 8F E C')
f4 = sequence('2C 2G3 2C 2=') 

v1 = join(f1,f1,f2,f2,f3,f3,f4) 

track('frere',1,
    onbar(1,v1),
    onbar(3,v1),
    onbar(5,v1))
```

</div>

<div class="copy-source"><button class="btn-clipboard btn-copy-source btn btn-sm btn-link" data-clipboard-target=".docs-content .highlight code" title="Copy to clipboard" aria-label="Copy to clipboard"><svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-copy" aria-hidden="true"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg></button></div>
