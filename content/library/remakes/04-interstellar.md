---
title: "04-interstellar"
menu:
  library:
    parent: remakes
weight: 40
---

<ul>
<li>Interstellar - Main Theme - Hans Zimmer</li>
<li><a href="https://musescore.com/user/19463776/scores/3816721">https://musescore.com/user/19463776/scores/3816721</a></li>
<li>Translated by Ernest Micklei using Melrōse</li>
<li>Analog Lab, Ambient Piano Pad</li>
</ul>

<div class="mel-source">

```javascript
// Interstellar - Main Theme - Hans Zimmer
//
// https://musescore.com/user/19463776/scores/3816721
//
// Translated by Ernest Micklei using Melrōse
// Analog Lab, Ambient Piano Pad


bpm(100)
biab(3)

e3 = sequence('e e e')
t1 = track("organ",1,
    onbar(1,join(dynamic('---',e3),dynamic('--',e3),dynamic('-',e3), repeat(43-3,e3) )))
t2 = track("melody",1
    ,onbar(5,sequence('(2.a 2.a3) (1.b 1.b3)')) // 7
    ,onbar(8,sequence('(a a3) (b b3) (c5 c) (b b3) (a a3) (b b3) (2.c5 2.c)'))
    ,onbar(11,sequence('(1.b 1.b3)')) // 12
    ,onbar(13,sequence('(2a 2a3) e5 (2.c5 2.c) (1.b 1.b3)')) // 16
    ,onbar(17,sequence('(a a3) e5 (c5 c) (b b3) (a a3) (b b3) (2.a 2.c)'))
    ,onbar(20,sequence('(1.b 1.b3)')) // 21 
    ,onbar(22,sequence('(a a3) e5 (c5 c) (b b3) (a a3) (b b3) (2.a 2.c) (2.d5 2.d) 2.e5')) // 26
    
    // 1/8
    ,onbar(27,sequence('(8a 2.a3) 8b 8c5 8b 8a 8c5 (8b 1.b3) 8a 8g 8a 8b 8g 8b 8a 8g 8a 8b 8g')) // 29
    ,onbar(30,fraction(8,sequence('(a a3) c5 (b b3) d5 (c5 c) a (b b3) g (a a3) c5 (b b3) d5'))) // 31
    ,onbar(32,sequence('(8c5 2.c) 8d5 8c5 8b 8a 8c5 (8b 1.b3) 8a 8g 8a 8b 8g 8b 8a 8g 8e5 8b 8g')) // 34
    
    // 1/16
    ,onbar(35,sequence('(16a 2.a) 16e 16b 16e 16c5 16e 16b 16e 16a 16e 16c5 16e'))
    ,onbar(36,sequence('(16b 1.a) 16e 16a 16e 16g 16e 16a 16e 16b 16e 16g 16e'))
    ,onbar(37,sequence('16b 16e 16a 16e 16g 16e 16e5 16e 16b 16e 16g 16e'))
    ,onbar(38,sequence('(16a3 a) 16e 16a 16e (16b3 b) 16e 16a 16e (16c c) 16e 16a 16e'))
    ,onbar(39,sequence('(16b3 b) 16e 16g 16e (16a3 a) 16e 16c 16e (16b3 b) 16e 16d5 16e'))
    ,onbar(40,sequence('(2.c 16c5) 16e 16d5 16e 16c5 16e 16b 16e 16a 16e 16c5 16e'))
    ,onbar(41,sequence('(1.b 16b) 16e 16a 16e 16g 16e 16a 16e 16b 16e 16g 16e'))
    ,onbar(42,sequence('16b 16e 16a 16e 16g 16e 16e5 16e 16b 16e 16g 16e'))
    ,onbar(43,sequence('(2.a 16a) 16e 16b 16e 16c5 16e 16b 16e 16e5 16e 16b 16e'))
    )

multitrack(t1,t2)
```

</div>

<div class="copy-source"><button class="btn-clipboard btn-copy-source btn btn-sm btn-link" data-clipboard-target=".docs-content .highlight code" title="Copy to clipboard" aria-label="Copy to clipboard"><svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-copy" aria-hidden="true"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg></button></div>
