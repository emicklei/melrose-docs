---
title: 04 Loop
menu:
  tutorial:
    parent: "basic"
---
## create a loop

```javascript
bpm(120)
s = sequence('c d e f g a b')
lp = loop(a)
```

A loop plays one or more musical objects repeatedly.
The tempo at which the notes of the objects are played is set using the `bpm` function.
The loop object must be assigned a the variable, here `lp` because the program needs a reference in order to `play` or `stop` the loop.

```javascript
play(lp)
stop(lp)
```

Using the editor with the plugin, it is much easier to play a loop using `cmd+3` and stop it with `cmd+5`.

## pitch loop

```javascript
i = interval(0,4,1)
p = pitch(i,sequence('c d e f g a b'))
l = loop(p,next(i))
```

An interval is a non-musical object that can generate integer numbers.
In this example, it will generate the numbers `0 1 2 3 4` and will repeat them in this order.
A pitch is a musical object modifier that changes the pitch of each note with the number of semitones.
In this example, the pitch of the notes of a sequence is increased by the current value of `i`.
On each loop entry, the pitched sequence is played and the interval is asked to go to the next value using `next(i)`.