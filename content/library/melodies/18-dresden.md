---
title: "18-dresden"
menu:
  library:
    parent: melodies
weight: 180
---

<ul>
<li>Dresden by Nil Zork</li>
<li>Analog Lab Pro: Wet Pluck, 80&#39;s FM EP</li>
</ul>

<div class="mel-source">

```javascript
// Dresden by Nil Zork
// Analog Lab Pro: Wet Pluck, 80's FM EP

s1 = sequence('.e 8e e_ e d_ e a')
s2 = sequence('.a 8a a_ a d_5 a g_')
s3 = sequence('.g_ 8g_ e g_ e_ e g_')
s4 = sequence('.a_ 8a_ g_ a e_ g_ e')
r = note('=')
j = stretch(0.5,join(s1,r,s2,r,s3,r,s4,r))
it = interval(-4,4,1)
lj = loop(pitch(it,j),next(it))
```

</div>

<div class="copy-source"><button class="btn-clipboard btn-copy-source btn btn-sm btn-link" data-clipboard-target=".docs-content .highlight code" title="Copy to clipboard" aria-label="Copy to clipboard"><svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-copy" aria-hidden="true"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg></button></div>
