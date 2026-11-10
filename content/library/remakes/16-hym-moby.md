---
title: "16-hym-moby"
menu:
  library:
    parent: remakes
weight: 160
---

<ul>
<li>Moby - Hym</li>
<li><a href="https://musescore.com/user/22928/scores/54657">https://musescore.com/user/22928/scores/54657</a></li>
<li>E flat major</li>
<li>Wavestate: set list 11, subtle thoughts</li>
<li>Analog Lab Pro: Emotional Piano</li>
<li>Garageband: FM Piano</li>
</ul>

<div class="mel-source">

```javascript
// Moby - Hym
// https://musescore.com/user/22928/scores/54657
// E flat major
// Wavestate: set list 11, subtle thoughts
// Analog Lab Pro: Emotional Piano
// Garageband: FM Piano


bpm(100)

s1=sequence('g5 e_5 c5 g c g c5 e_5')
s2=sequence('b_5 g5 e_5 b_ e_ b_ e_5 g5')
s3=sequence('c6 a_5 f5 c5 f c5 f5 a_5')
s4=sequence('E_6 C6 G5 E_5 C5 E_5 G5 C6')

p1=repeat(4,fraction(16,s1))
p2=replace(p1,s1,s2)
p3=replace(p1,s1,s3)
p4=replace(p1,s1,s4)

t1 = join(p1,p2,p3,p4,note("="))

repeat(4,t1)
```

</div>

<div class="copy-source"><button class="btn-clipboard btn-copy-source btn btn-sm btn-link" data-clipboard-target=".docs-content .highlight code" title="Copy to clipboard" aria-label="Copy to clipboard"><svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-copy" aria-hidden="true"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg></button></div>

Copy the script and try it using <a href="https://play.melrōse.org" target="_blank">Melrōse Playground</a>.
