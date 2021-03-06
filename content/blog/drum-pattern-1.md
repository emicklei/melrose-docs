---
title: "Drum Pattern"
description: "Using notemap functions to create a drum beat."
lead: "Using notemap functions to create a drum beat."
date: 2021-03-06T16:53:58+01:00
lastmod: 2021-03-06T16:53:58+01:00
draft: false
weight: 50 
contributors: []
---

This script is an example that uses the Melrōse language to create your own drum beats.
Line by line, I will explain how it is composed.
 
{{< highlight javascript "linenos=table,linenostart=1" >}}

bpm(85)

// https://usermanuals.finalemusic.com/SongWriter2012Win/Content/PercussionMaps.htm

open = midi(16,46,60)
close = midi(16,42,60)
clap = midi(16,39,60)
kick = midi(16,36,80)

d1 = notemap(' 3 10 11',open)
d2 = notemap('4 12',close)
d3 = notemap('5 13',clap)
d4 = notemap('1 2 4 7 9 10 12 15',kick)
d5 = notemap('16',note('16=')) // 1/16 rest at the end to make 16 notes in the set
set = merge(d1,d2,d3,d4,d5)
lp_set = loop(set)

{{< / highlight >}}

<a href="https://storage.googleapis.com/downloads.ernestmicklei.com/melrose/melrose_drum_pattern-1.aif">🎶 Play</a>


### line 1
Using the bpm function you change the default from 120 beats-per-minute to 85.

### line 5 .. 8
Using the General MIDI mapping (see url on line 3),  4 variables with notes are created with the `midi` function.  Parameter 16 refers to the fraction of the note,  1/16. The second parameter is the MIDI number, the kick is 36. The third parameter is the velocity (loudness) and is set a bit lower than the default (72).

### line 10 .. 13
Using the notemap function we create a sequence of notes in which for each number the note is placed. On index 3 , 10 and 11 (sequences start at 1), the open note is placed. On all other indices (1,2,...) a rest note ( '=' ) is placed. 
For each note (open,close,clap,kick) a pattern is created using notemap.

The indices in this example were found using [Ableton learning music](https://learningmusic.ableton.com/make-beats/make-beats.html). 

An alternative way to describe the indices in a `notemap` is using dots (.) and bangs (!).

```javascript
d1 = notemap('..!......!!.....') // 16 characters
```

Which gives a better visual feedback of the pattern.

### line 14
The drum set is created by merging all the notemaps into one sequence. The first parameter 16 of the `merge` (notemerge is the old name) function is needed to specify the total number of notes in the set. Again, unspecified indices will have the rest note.

### line 15
Finally, on this line we create a loop using the set as its parameter. You can play this loop by using cmd+3 in the Editor (*)

### Modify while play

While playing the loop, you can change the notes and the notemap in the Editor and evaluate each line to hear the effect.
Below is a snapshot of a piano roll created from playing this loop using a DAW (e.g. Logic).
 
![drum pattern pianoroll](https://storage.googleapis.com/downloads.ernestmicklei.com/melrose/melrose-drum-pattern1-pianoroll.png) 
 

(*) Editor refers to the Melrōse Plugin for Visual Studio Code which is connected to the program melrose which send MIDI message to your Digital Audio Workstation (DAW).
