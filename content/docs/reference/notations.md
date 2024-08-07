
---
title: "Notation"
description: "Musical object notations."
lead: "Musical object notations."
date: 2021-03-05T15:21:01+02:00
lastmod: 2021-03-09T15:21:01+02:00
draft: false
images: []
menu: 
  docs:
    parent: "reference"
weight: 10
toc: true
---


## Note
<a name="note"></a>

Format: `(fraction)(dot)(name|=)(accidental)(dynamic)`

Note names are case-insentivive but are always displayed in uppercase.

| Notation | Alternative | Description
|----------|-------------|-------------
| C4       | c          | quarter C, octave 4
| 2e5      | 2E5        | Halftone (2 x ¼), E octave 5
| 1c       | 1C         | Full tone C, octave 4
| f#       | F#, F♯     | F sharp
| g_       | G_, G♭     | G flat
| .g       | .G         | duration fraction x 1.5 = 3/8
| =        | =          | quarter rest
| 2=       |            | half rest
| 1=       | 1=         | full rest
| d+       | D+         | quarter D, octave 4, MezzoForte
| 16.e#--  | 16.E#--    | sixteenth, E sharp, fraction x 1.5, Piano
| 32f      | 32F        | thirtytwo F
 
```javascript
n = note('c#5')
```

## Note dynamics
<a name="dynamics"></a>

| Notation    | Description
|-------------|---
| \-\-\-\-\-    |Pianissississimo (pppp)
| \-\-\-\-      |Pianississimo (ppp)
| \-\-\-        |Pianissimo (pp)
| \-\-          |Piano (p)
| -             |MezzoPiano (mp)
| o (not 0)     |Normal (character is optional)
| +             |MezzoForte (mf)
| ++            |Forte (f)
| +++           |Fortissimo (ff)
| ++++          |Fortississimo (fff)
| +++++         |Fortissississimo (ffff)

```javascript
n = note('e++')
```

## Sequence
<a name="sequence"></a>

| Notation    | Description
|-------------|---
| C D E F        | 4 quarter tones
| (8C E) (d5 f5) | 2 doublets; first doublet has an eight length, second is a quarter
| (1C E G)    | C Chord; whole length

```javascript
doremi = sequence('c d e')
```

## Pedal control

Usable in `sequence` or `note`.

| Notation | Description
|----------|-------------
| >        | sustain pedal `down`
| <        | sustain pedal `up`
| ^        | sustain pedal `up` and immediately `down`


## Chord
<a name="chord"></a>

| Notation    | Description
|-------------|---
| c#5/m/2     | C sharp triad, Octave 5, Minor, 2nd inversion
| a/7         | A Dominant seventh chord
| e/maj7      | E Major seventh chord
| g/m7        | G minor seventh chord
| 1=          | No chord, a whole rest note
| d/dim       | D diminished triad
| d/o         | D diminished triad
| f/dim7/1    | F diminished seventh, 1st inversion
| c/aug       | C augmented triad
| e/+         | E augmented triad
| b_/+7       | B flat augmented seventh

```javascript
b7 = chord('b/7')
```

## Scale
<a name="scale"></a>

| Notation    | Description
|-------------|---
| c5          | C major scale, Octave 5
| e/m         | E natural minor scale, Octave 4
| g/maj7      | G major 7 scale, Octave 4

```javascript
sf = scale('f')
```

## Chord Progression 
<a name="progression"></a>

| Notation    | Alternative | Description
|-------------|--------|--
| I           | i      | first chord in scale ; if scale is "C" then sequence is "(C E G)"
| V7          | v7     | (G B D5 F5)
| Imaj7       | imaj7  | (C E G B)
| viidim      | VIIdim | (B D5 F5)

```javascript
p = progression('c', 'ii V I') // Major C scale, (D F A) (G B D5) (C E G)
```

## Chord Sequence 
<a name="chordsequence"></a>

| Notation    | Description
|-------------|---
| c/m d/m     | C minor followed by a D minor
| (c3 c5)     | C major, Octave 3 together with a C major, Octave 5
| e =         | E major followed by a quarter rest note

```javascript
cs = chordsequence('c f g')
```