---
title: "10-telios"
menu:
  library:
    parent: melodies
weight: 100
---

<ul>
<li>Telios by Nil Zork</li>
</ul>

<div class="mel-source">

```javascript
// Telios by Nil Zork
//

bpm(140)

s1 = sequence('c3# g3# c# e g# c5#')
s2 = sequence('e3 b3 e g# b e5')
s3 = sequence('f3# c# f# a c#5 f#5')
s4 = sequence('b2 f3# b3 d# f# b')

loop(fraction(8,join(s1,s3,s2,s4)))
```

</div>

<div class="copy-source"><button class="btn-clipboard btn-copy-source btn btn-sm btn-link" data-clipboard-target=".docs-content .highlight code" title="Copy to clipboard" aria-label="Copy to clipboard"><svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-copy" aria-hidden="true"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg></button></div>
