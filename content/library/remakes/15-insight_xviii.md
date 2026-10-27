---
title: "15-insight_xviii"
menu:
  library:
    parent: remakes
weight: 150
---

<ul>
<li>Insight XVIII by Jean Michel Blais</li>
<li>Garageband: Massive Trance Pad</li>
</ul>

<div class="mel-source">

```javascript
// Insight XVIII by Jean Michel Blais
// Garageband: Massive Trance Pad

bpm(70)

t1 = seq('(C#3 G#3 C# E G# E5)') // chord('c#/m')
t2 = seq('(D#3 B3 D# F# B F#5)')

t3 = seq('(E3 B3 E G# B G#5)')
t4 = seq('(F#3 C# F# A C#5 A5)')
t5 = seq('(B2 F#3 B3 D# F# D#5)')
t6 = seq('(A2 E3 A3 C# E C#5)')
t7 = seq('(G#2 D#3 G#3 C D# C5)')

tall = join(t1, t2, t3, stretch(2,t4), stretch(2,t5), t1, t5, t6, stretch(2,t7))
lp_tall = loop(tall)
```

</div>

<div class="copy-source"><button class="btn-clipboard btn-copy-source btn btn-sm btn-link" data-clipboard-target=".docs-content .highlight code" title="Copy to clipboard" aria-label="Copy to clipboard"><svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-copy" aria-hidden="true"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg></button></div>
