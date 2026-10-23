---
title: "02-tubular-bells-multi"
menu:
  library:
    parent: remakes
weight: 20
---

<ul>
<li>Tubular Bells by Mike Oldfield (Opening Theme in 15/8)</li>
<li>Layered Multi-Track Arrangement</li>
<li>To play both the piano lead and the tubular bell accent together</li>
<li>on separate MIDI channels:</li>
</ul>

<div class="mel-source">

```javascript
// Tubular Bells by Mike Oldfield (Opening Theme in 15/8)
// Layered Multi-Track Arrangement
// To play both the piano lead and the tubular bell accent together 
// on separate MIDI channels:

bpm(144)

// Piano Lead (MIDI Channel 1)
motif1 = fraction(8,sequence('e a e b e g a e c5 e d5 e b c5'))
motif2 = fraction(8,sequence('e a e b e g a e c5 e d5 e b c5 e b'))
piano_lead = channel(1, join(motif1, motif2))

// Low tubular bell strikes on the downbeats (MIDI Channel 2)
// GM Program 15 corresponds to Tubular Bells
chimes = channel(2, sequence('1e3 1= 1d3 1='))

// Create synchronized loops
loop_piano = loop(piano_lead)
loop_chimes = loop(chimes)

// Play both loops together
sync(loop_piano,loop_chimes)
```

</div>

<div class="copy-source"><button class="btn-clipboard btn-copy-source btn btn-sm btn-link" data-clipboard-target=".docs-content .highlight code" title="Copy to clipboard" aria-label="Copy to clipboard"><svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-copy" aria-hidden="true"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg></button></div>
