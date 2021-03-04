---
title: 02 Sequence
menu:
  docs:
    parent: "advanced"
---

## create a sequence

```javascript
sequence('c d (e f)')
```

The expression uses the function `sequence` and takes the string argument `'c d (e f)'`.
This argument uses the notation for a note.
The notes inside brackets, here E and F, will be played as the same time.
Another, more complex example is:

```javascript
sequence('(8c e g) = (2c++ e++ g++)')
```

The first group `(8c e g)` has fraction 1/8 the second 1/2. 
Only the first note should set the fraction of the group.

### sustain pedal

In a sequence, you can control the position of the sustain pedal.
To push the pedal, you use the character `>`.
To release the pedal, you use the character `<`.
The character `^` is short for up and down.

```javascript
sequence('> c d e ^ e d c <')
```