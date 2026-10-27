---
title: "03-sakamoto-harry"
menu:
  library:
    parent: remakes
weight: 30
---

<ul>
<li>Harry to Hospital - R Sakamoto</li>
<li>Logic Pro X, Saturn Lead</li>
<li>Analog Lab, Ambient Piano Pad</li>
</ul>

<div class="mel-source">

```javascript
// Harry to Hospital - R Sakamoto
// Logic Pro X, Saturn Lead
// Analog Lab, Ambient Piano Pad

bpm(70)
s1 = stretch(0.25,sequence('> (f#3 c#5) c# f# 2a a f# c# <'))
s2 = stretch(0.25,sequence('> (e3 g# c#5) c# e 2g# g# e c# <'))
join(repeat(2,s1),repeat(2,s2))

// separate lanes
l2 = fraction(8,sequence('f#3 = f#3- = f#3-- = f#3--- ='))
l3 = pitch(-2,l2)
l4 = pitch(-4,l2)

l234 = join(l2,l3,l4,l3)

r2 = repeat(2,stretch(0.25,sequence('c#5 c# f# 2a a f#- c#--')))
r3 = repeat(2,stretch(0.25,sequence('c#5 c# e 2g# g# e- c#--')))
r4 = repeat(2,stretch(0.25,sequence('c#5 c# e 2f# f# e- c#--')))
r5 = repeat(2,stretch(0.25,sequence('b   b3 e 2g# g# e- b3--')))

r2345 = join(r2,r3,r4,r5)

// sync(l234, r2345)

m = merge(l234,r2345)

lp_harry = loop(m)
```

</div>

<div class="copy-source"><button class="btn-clipboard btn-copy-source btn btn-sm btn-link" data-clipboard-target=".docs-content .highlight code" title="Copy to clipboard" aria-label="Copy to clipboard"><svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-copy" aria-hidden="true"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg></button></div>
