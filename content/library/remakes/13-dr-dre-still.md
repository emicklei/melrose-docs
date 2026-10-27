---
title: "13-dr-dre-still"
menu:
  library:
    parent: remakes
weight: 130
---

<ul>
<li>Dr. Dre ft. Snoop Dogg – &#34;Still D.R.E.&#34;</li>
<li>Garageband: FM Piano</li>
</ul>

<div class="mel-source">

```javascript
// Dr. Dre ft. Snoop Dogg – "Still D.R.E."
// Garageband: FM Piano

bpm(93)

// 8 syncopated hits on A minor, followed by 3 on E minor and 5 on B minor
hit_am = sequence('8(a4 c5 e5)')
hit_em = sequence('8(g4 b4 e5)')
hit_bm = sequence('8(f#4 b4 d5)')

bar = join(
    hit_am, hit_am, hit_am, hit_am, hit_am, hit_am, hit_am, hit_am,
    hit_em, hit_em, hit_em,
    hit_bm, hit_bm, hit_bm, hit_bm, hit_bm
)

// Bass groove underneath
bass = sequence('4.a2 8= 4.a2 8= 4.e2 8= 4.b2 8=')

l1 = loop(channel(1, bar))
l2 = loop(channel(2, bass))

sync(l1, l2)
```

</div>

<div class="copy-source"><button class="btn-clipboard btn-copy-source btn btn-sm btn-link" data-clipboard-target=".docs-content .highlight code" title="Copy to clipboard" aria-label="Copy to clipboard"><svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-copy" aria-hidden="true"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg></button></div>
