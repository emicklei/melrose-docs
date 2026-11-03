---
title: "21-kamicros"
menu:
  library:
    parent: melodies
weight: 210
---

<ul>
<li>Kamicros by Nil Zork</li>
<li>Ableton: FM Piano5</li>
<li>Multi channel</li>
</ul>

<div class="mel-source">

```javascript
// Kamicros by Nil Zork
// Ableton: FM Piano5
// Multi channel

bpm(120)

s1 = sequence('> g b_ e_5 e_ b_ e_5 <')
s2 = sequence('> g b_ d5 d b_ d5 <')
s3 = sequence('> c g c5  e_ g  c5 <')
s4 = sequence('> f c5 f5 b_  c5 f5 <')

j1 = join(
    repeat(2,fraction(8,s1)),
    repeat(2,fraction(8,s2)),
    repeat(2,fraction(8,s3)),
    repeat(2,fraction(8,s4))
    )

loop(channel(1,j1))

// set('midi.out',1)

b1 = repeat(2,octave(-1,sequence('.g .e_')))
b2 = repeat(2,octave(-1,sequence('.g .d'))) 
b3 = repeat(2,octave(-1,sequence('.c .e_')))
b4 = repeat(2,octave(-1,sequence('.f .b_')))

j2 = join(b1,b2,b3,b4)

loop(channel(2,j2))
```

</div>

<div class="copy-source"><button class="btn-clipboard btn-copy-source btn btn-sm btn-link" data-clipboard-target=".docs-content .highlight code" title="Copy to clipboard" aria-label="Copy to clipboard"><svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-copy" aria-hidden="true"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg></button></div>
