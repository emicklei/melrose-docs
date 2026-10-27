---
title: "07-domeka"
menu:
  library:
    parent: melodies
weight: 70
---

<ul>
<li>Domeka by Nil Zork</li>
<li>Logix Pro, Ethereal Mallets</li>
<li>Multi channel</li>
</ul>

<div class="mel-source">

```javascript
// Domeka by Nil Zork
// Logix Pro, Ethereal Mallets
// Multi channel

bpm(100)
p = iterator(6,9,4,8)
a2 = octavemap('1:0,3:0,2:1,1:1,3:1,1:1,2:1,3:0',ungroup(pitch(p,chord('8c3/m'))))
lp_a2 = loop(a2, next(p))

track("Ethereal",1,
    onbar(1,repeat(4,a2,next(p))),
    onbar(5,repeat(4,a2,next(p))))

p = iterator(0,4,7,12)    

kick = channel(2,sequence('8c2 8c#2 8= 8= 8d2 8d#2 8= 8='))
lp_kick = loop(kick)
sync(lp_a2, lp_kick)
```

</div>

<div class="copy-source"><button class="btn-clipboard btn-copy-source btn btn-sm btn-link" data-clipboard-target=".docs-content .highlight code" title="Copy to clipboard" aria-label="Copy to clipboard"><svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-copy" aria-hidden="true"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg></button></div>
