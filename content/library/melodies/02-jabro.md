---
title: "02-jabro"
menu:
  library:
    parent: melodies
weight: 20
---

<ul>
<li>Jabro by Nil Zork</li>
<li>Ableton: June O Pad</li>
<li>Wavestation: Five On It Bells</li>
<li>Analog Lab V: Bouncy Bells</li>
</ul>

<div class="mel-source">

```javascript
 // Jabro by Nil Zork
 // Ableton: June O Pad
 // Wavestation: Five On It Bells
 // Analog Lab V: Bouncy Bells

bpm(120)
s1 = sequence('a2 e3 c# e3 a2 e3 c#')
s2 = sequence('b2 f#3 d f#3 b2 f#3 d')
s3 = sequence('d3 a3 f# a3 d3 a3 f#')
s4 = sequence('c#3 g#3 e g#3 c#3 g#3 e')
cstep = note('c#3')
dstep = note('d3')

j1 = join(s1,cstep ,s2,cstep ,s3,dstep, s4)
j2 = fraction(8,join(j1,cstep,j1,cstep))

lp_j2 = loop(dynamic('--',j2))

// export('19-jabro.mid', lp_j2)
```

</div>

<div class="copy-source"><button class="btn-clipboard btn-copy-source btn btn-sm btn-link" data-clipboard-target=".docs-content .highlight code" title="Copy to clipboard" aria-label="Copy to clipboard"><svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-copy" aria-hidden="true"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg></button></div>
