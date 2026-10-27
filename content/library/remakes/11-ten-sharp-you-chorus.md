---
title: "11-ten-sharp-you-chorus"
menu:
  library:
    parent: remakes
weight: 110
---

<ul>
<li>Full Multi-Track Chorus (&#34;You, you were always on my mind...&#34;)</li>
<li>Multi channel</li>
<li>1: vocal</li>
<li>2: keys</li>
<li>3: bass</li>
</ul>

<div class="mel-source">

```javascript
// Full Multi-Track Chorus ("You, you were always on my mind...")
// Multi channel
// 1: vocal
// 2: keys
// 3: bass


// -------------------------------------------------------------
// Ten Sharp - "You" (Chorus Section)
// -------------------------------------------------------------

bpm(87)

// --- 1. Vocal Melody Line ---
// "You... you were always on my mind"
m_line1 = sequence('2f#4 8= 8f#4 8e4 8d4 8e4 8f#4 4d4 2=')

// "You... you're the one I'm living for"
m_line2 = sequence('2g4  8= 8g4  8f#4 8e4 8f#4 8g4 4e4 2=')

// "You... you're my everlasting fire"
m_line3 = sequence('2a4  8= 8a4  8g4  8f#4 8g4  8a4 4f#4 2=')

// "Ooh... you're my always shining star"
m_line4 = sequence('4b4 4a4 4g4 4f#4 4e4 4d4 2d4')

vocal_chorus = join(m_line1, m_line2, m_line3, m_line4)

// --- 2. Piano & Pad Chords ---
ch_d  = sequence('2(d4 f#4 a4) 2(d4 f#4 a4)')
ch_g  = sequence('2(d4 g4 b4)  2(d4 g4 b4)')
ch_em = sequence('2(e4 g4 b4)  2(e4 g4 b4)')
ch_a  = sequence('2(e4 a4 c#5) 2(e4 a4 c#5)')

ch_bar7 = sequence('2(e4 g4 b4)  2(e4 a4 c#5)')
ch_bar8 = sequence('1(d4 f#4 a4)')

chorus_chords = join(ch_d, ch_g, ch_em, ch_a, ch_d, ch_g, ch_bar7, ch_bar8)

// --- 3. Chorus Bass Line ---
c_b1 = sequence('2d2 4d2 4d2')
c_b2 = sequence('2g2 4g2 4g2')
c_b3 = sequence('2e2 4e2 4e2')
c_b4 = sequence('2a2 4a2 4a2')
c_b5 = sequence('2d2 4d2 4d2')
c_b6 = sequence('2g2 4g2 4g2')
c_b7 = sequence('2e2 2a2')
c_b8 = sequence('1d2')

chorus_bass = join(c_b1, c_b2, c_b3, c_b4, c_b5, c_b6, c_b7, c_b8)

// --- Assign MIDI Channels ---
lead_ch  = channel(1, vocal_chorus)
keys_ch  = channel(2, chorus_chords)
bass_ch2 = channel(3, chorus_bass)

// Play the chorus
sync(lead_ch, keys_ch, bass_ch2)
```

</div>

<div class="copy-source"><button class="btn-clipboard btn-copy-source btn btn-sm btn-link" data-clipboard-target=".docs-content .highlight code" title="Copy to clipboard" aria-label="Copy to clipboard"><svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-copy" aria-hidden="true"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg></button></div>
