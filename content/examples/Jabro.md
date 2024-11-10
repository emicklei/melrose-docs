---
title: "Jabro"
date: 2022-03-13
draft: false
---

 - Ableton: June O Pad
 - WS: Five On It Bells
 - Analog Lab V: Bouncy Bells

```javascript
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
