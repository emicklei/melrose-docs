---
title: "04-riboluta"
menu:
  library:
    parent: melodies
weight: 40
---

<ul>
<li>Riboluta by Nil Zork</li>
<li>Logic: Micro Pulse</li>
<li>Ableton: June O Pad</li>
<li>Garageband: Percussive Square Lead</li>
</ul>

<div class="mel-source">

```javascript
// Riboluta by Nil Zork
// Logic: Micro Pulse
// Ableton: June O Pad
// Garageband: Percussive Square Lead

bpm(120)

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

</div>

<div class="copy-source"><button class="btn-clipboard btn-copy-source btn btn-sm btn-link" data-clipboard-target=".docs-content .highlight code" title="Copy to clipboard" aria-label="Copy to clipboard"><svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-copy" aria-hidden="true"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg></button></div>
