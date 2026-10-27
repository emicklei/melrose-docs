---
title: "10-ten-sharp-you"
menu:
  library:
    parent: remakes
weight: 100
---

<ul>
<li>Ten Sharp - &#34;You&#34; (Intro / Verse Theme)</li>
<li>Multi channel</li>
<li>1: piano</li>
<li>2: bass</li>
</ul>

<div class="mel-source">

```javascript
// Ten Sharp - "You" (Intro / Verse Theme)
// Multi channel
// 1: piano
// 2: bass

bpm(87)

// --- Right-hand arpeggiated piano pattern (16th notes) ---

// Bar 1: B minor (Bm)
p_bm = sequence('16b3 16f#4 16b4 16d5 16f#4 16b4 16d5 16f#4 16b3 16f#4 16b4 16d5 16f#4 16b4 16d5 16f#4')

// Bar 2: E major (E)
p_e  = sequence('16e3 16g#4 16b4 16e5 16g#4 16b4 16e5 16g#4 16e3 16g#4 16b4 16e5 16g#4 16b4 16e5 16g#4')

// Bar 3: E minor (Em) to A major (A)
p_em = sequence('16e3 16g4 16b4 16e5 16g4 16b4 16e5 16g4')
p_a  = sequence('16a3 16e4 16a4 16c#5 16e4 16a4 16c#5 16e4')
p_bar3 = join(p_em, p_a)

// Bar 4: D major (D) to F# major (F#)
p_d  = sequence('16d3 16f#4 16a4 16d5 16f#4 16a4 16d5 16f#4')
p_fs = sequence('16f#3 16a#4 16c#5 16f#5 16a#4 16c#5 16f#5 16a#4')
p_bar4 = join(p_d, p_fs)

// Full 4-bar piano loop
piano_riff = join(p_bm, p_e, p_bar3, p_bar4)

// --- Bass Line ---
b_bm = sequence('4b2 8= 8b2 4b2 4b2')
b_e  = sequence('4e2 8= 8e2 4e2 4e2')
b_ea = sequence('2e2 2a2')
b_df = sequence('2d2 2f#2')

bass_riff = join(b_bm, b_e, b_ea, b_df)

// Assign channels (Channel 1: Bright Piano, Channel 2: Bass)
piano_ch = channel(1, loop(piano_riff))
bass_ch  = channel(2, loop(bass_riff))

// Play the intro groove
sync(piano_ch, bass_ch)
```

</div>

<div class="copy-source"><button class="btn-clipboard btn-copy-source btn btn-sm btn-link" data-clipboard-target=".docs-content .highlight code" title="Copy to clipboard" aria-label="Copy to clipboard"><svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-copy" aria-hidden="true"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg></button></div>
