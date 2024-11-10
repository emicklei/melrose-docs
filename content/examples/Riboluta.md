---
title: "Riboluta"
date: 2021-10-04
draft: false
---
- Logic: Micro Pulse
- Ableton: June O Padbpm(120)
- Garageband: Percussive Square Lead

```javascript
y = sequence('f#2 c#3 f#3 a3 c# f# = ')
p = resequence('3 4 2 5 1 6 2 5', y)
f = fraction(8, p)
jf = join(repeat(2,f),
	repeat(2,pitch(1,f)), 
	repeat(2,pitch(-2,f)), 
	repeat(2,pitch(3,f)))
loop(jf)

// lp_jp = loop(channel(1,jp))
tr1 = track("lead",1,jp,jp,jp,jp)

// 2nd voice, Deep Sub Bass
p2 = resequence('3 7 2 7 1 7 2 7', y) // 7 = rest note
jp2 = replace(jp,p,p2)

// lp_jp2 = loop(channel(4,jp2))
tr2 = track("2nd",2,
	onbar(bars(jp)+1,jp2),
	onbar(bars(jp)*2+1,jp2),
	onbar(bars(jp)*3+1,jp2))



// 3rd voice, Black Whisper
p3 = resequence('5 7 7 7 6 7 7 7', y)
jp3 = replace(jp,p,p3)

// lp_jp3 = loop(channel(2,jp3))
tr3 = track("3nd",3,
	onbar(bars(jp)*2+1,jp3),
	onbar(bars(jp)*3+1,jp3)) 
 


// 4th voice, Liverpool drumset
// y4 = sequence('c2 g#2 g#2 b2 c2 g#2 g#2 b2')
y4 = sequence('= g#2 = g#2 = g#2 = g#2')
tr4 = track("4th",4,
	onbar(bars(jp)*3+1,repeat(8,y4)))

m = multitrack(tr1,tr2,tr3,tr4)

// export("riboluta-v1",m)
```
