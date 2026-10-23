---
title: "01-tubular-bells"
menu:
  library:
    parent: remakes
weight: 10
---

<ul>
<li>-----------------------------------------------------------</li>
<li>Mike Oldfield - Tubular Bells (Opening Theme in 15/8)</li>
<li>-----------------------------------------------------------</li>
</ul>

<div class="mel-source">

```javascript
// -----------------------------------------------------------
// Mike Oldfield - Tubular Bells (Opening Theme in 15/8)
// -----------------------------------------------------------

// 1. Set the tempo (approx. 144-150 BPM for eighth notes)
bpm(144)

// 2. Define the first phrase: 7/8 measure (14 eighth notes)
// '8' sets the duration fraction to an eighth note (1/8).
// Octave 4 is default, 'c5' and 'd5' jump to the 5th octave.
part1 = fraction(8,sequence('e a e b e g a e c5 e d5 e b c5'))

// 3. Define the second phrase: 8/8 measure (16 eighth notes)
// Ends with the extra two eighth notes (e, b) that complete the 15/8 cycle
part2 = fraction(8,sequence('e a e b e g a e c5 e d5 e b c5 e b'))

// 4. Combine both phrases into the full 30-note 15/8 motif
tubular_theme = join(part1, part2)

// 5. (Optional) Define the subtle low bass accompaniment (A minor -> G)
// 2-beat half notes (2) and whole rests (=) to anchor the 15/8 cycle
bass_riff = sequence('2a2 2a2 2g2 2g2 1=')

// 6. Play the motif in a continuous loop
loop(tubular_theme)
loop(bass_riff)
```

</div>

<div class="copy-source"><button class="btn-clipboard btn-copy-source btn btn-sm btn-link" data-clipboard-target=".docs-content .highlight code" title="Copy to clipboard" aria-label="Copy to clipboard"><svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-copy" aria-hidden="true"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg></button></div>
