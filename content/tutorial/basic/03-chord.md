---
title: 03 Chord
menu:
  tutorial:
    parent: "basic"
---

## create a chord

```javascript
chord('C#/m')
```

A chord is a special sequence of notes.
This expression uses the function `chord` and takes the string argument `'C#/m'`.
This argument uses the notation for a chord which uses a note notation for the base.
The `m` indicates a minor chord.
You can also specify an inversion:

```javascript
chord('2c#5/m/2')
```

half duration, C sharp, octave 5, minor, 2nd inversion