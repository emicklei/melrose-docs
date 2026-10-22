---
title: "05-waimuso"
menu:
  library:
    parent: melodies
weight: 50
---

<ul>
<li>Waimuso by Nil Zork</li>
<li>Logic Pro X</li>
<li>Multiple loops with different instruments, each on another channel.</li>
</ul>

<div class="mel-source">

```javascript
// Waimuso by Nil Zork
// Logic Pro X
// Multiple loops with different instruments, each on another channel.

bpm(120)

s1 = sequence('1c3--- 1e3--- 1a2--- 1d3---')
l1 = loop(channel(2,s1)) // end(l1)   --- black whisper


l2 = loop(channel(4,s1)) // end(l2)   --- subby bass


s2 = sequence('= 8a3 8b3 2c-  = 8c 8d 2e-  = 8d 8e 2f-   = 8c 8b3 2a3-') 
l3 = loop(channel(1,s2)) // end(l3)   --- grand piano + pad
 

s3 = repeat(4,sequence('= g#2+'))
l4 = loop(channel(3,s3)) // end(l4)   --- SoCal
```

</div>

<div class="copy-source"><button class="btn-clipboard btn-copy-source btn btn-sm btn-link" data-clipboard-target=".docs-content .highlight code" title="Copy to clipboard" aria-label="Copy to clipboard"><svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-copy" aria-hidden="true"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg></button></div>
