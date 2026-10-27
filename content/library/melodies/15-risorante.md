---
title: "15-risorante"
menu:
  library:
    parent: melodies
weight: 150
---

<ul>
<li>Analog Lab,</li>
</ul>

<div class="mel-source">

```javascript
// Analog Lab, 
bpm(120)
r1 = sequence('f d_5 d_5 c5 b_ a b_ c5')
r2 = sequence('f e_5 e_5 d_5 c5 b_ c5 d_5')
r3 = sequence('a_ f5 f5 e_5 d_5 c5 d_5 e_5')
r4 = sequence('a_ e_5 e_5 d_5 c5 b_ c5 d_5')
loop(fraction(8,join(r1,r2,r3,r4))) 



// Variation1:
// Syncopated Funk / Rhythmic Phrasing (Mixed Note Lengths & Rests)

bpm(120)

// 4 = quarter note, 8 = eighth note, 16 = sixteenth note, = = rest
v1 = sequence('4f 16d_5 16d_5 8c5 8= 16b_ 16a 8b_ 4c5')
v2 = sequence('4f 16e_5 16e_5 8d_5 8c5 8b_ 16c5 16d_5 4=')
v3 = sequence('8a_ 8= 16f5 16f5 8e_5 8d_5 16c5 16d_5 4e_5')
v4 = sequence('8a_ 16e_5 16e_5 8d_5 8c5 16b_ 16c5 4d_5 8=')

loop(join(v1, v2, v3, v4))


// Variation 2:
// Grouping roots and top notes into simultaneous intervals/chords

h1 = sequence('(f c5) 8d_5 8d_5 8c5 (b_ d_5) 8a 8b_ (f c5)')
h2 = sequence('(f d_5) 8e_5 8e_5 8d_5 (c5 e_5) 8b_ 8c5 (f d_5)')
h3 = sequence('(a_ c5) 8f5 8f5 8e_5 (d_5 f5) 8c5 8d_5 (a_ e_5)')
h4 = sequence('(a_ c5) 8e_5 8e_5 8d_5 (b_ d_5) 8b_ 8c5 (f c5)')

loop(fraction(8, join(h1, h2, h3, h4)))


// Variation 3:
// Algorithmic Grouping with resequence (3+3+2 Polyrhythmic Feel)

bpm(120)

r1 = sequence('f d_5 d_5 c5 b_ a b_ c5')
r2 = sequence('f e_5 e_5 d_5 c5 b_ c5 d_5')
r3 = sequence('a_ f5 f5 e_5 d_5 c5 d_5 e_5')
r4 = sequence('a_ e_5 e_5 d_5 c5 b_ c5 d_5')

// Re-group indices: accents on 1, 4, 7 (3+3+2 group feel)
accent_pattern = '1 2 3 4 5 6 7 8 1 3 5 7'

g1 = resequence(accent_pattern, r1)
g2 = resequence(accent_pattern, r2)
g3 = resequence(accent_pattern, r3)
g4 = resequence(accent_pattern, r4)

// Play at 16th note speed for a driving arpeggiated drive
loop(fraction(16, join(g1, g2, g3, g4)))


// Variation 4:
// Dynamic Contrast (Half-Time Call & Double-Time Response)

bpm(120)

r1 = sequence('f d_5 d_5 c5 b_ a b_ c5')
r2 = sequence('f e_5 e_5 d_5 c5 b_ c5 d_5')
r3 = sequence('a_ f5 f5 e_5 d_5 c5 d_5 e_5')
r4 = sequence('a_ e_5 e_5 d_5 c5 b_ c5 d_5')

// r1 & r3 in half-time (quarter notes: fraction 4)
// r2 & r4 in double-time (sixteenth notes: fraction 16)
slow_part = fraction(4, join(r1, r3))
fast_part = fraction(16, join(r2, r4))

loop(join(slow_part, fast_part))

// Variation 5:
// Polyphonic Layering (Grouped Bassline + Staccato Melody)

bpm(120)

// Fast 16th-note staccato melody
lead = fraction(16, join(
    sequence('f d_5 d_5 c5 b_ a b_ c5'),
    sequence('f e_5 e_5 d_5 c5 b_ c5 d_5'),
    sequence('a_ f5 f5 e_5 d_5 c5 d_5 e_5'),
    sequence('a_ e_5 e_5 d_5 c5 b_ c5 d_5')))

// Sustained half-note and whole-note bassline anchoring each phrase
bass = sequence('1f2  1a_2')

// Merge both streams together into a single loop
loop(merge(lead, bass))
```

</div>

<div class="copy-source"><button class="btn-clipboard btn-copy-source btn btn-sm btn-link" data-clipboard-target=".docs-content .highlight code" title="Copy to clipboard" aria-label="Copy to clipboard"><svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-copy" aria-hidden="true"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg></button></div>
