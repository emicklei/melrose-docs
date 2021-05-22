---
title: "Language"
description: "All language functions grouped by creation,compostion and audio control."
lead: "All language functions grouped by creation,compostion and audio control."
date: 2021-03-05T12:21:01+02:00
lastmod: 2021-03-05T12:21:01+02:00
draft: false
images: []
menu: 
  docs:
    parent: "reference"
weight: 20
toc: true
---

## Structure

### Expressions

Musical objects are created, composed and played using the <strong>melrõse</strong> tool by evaluating expressions.
Expressions use any of the predefined functions (creation,composition,audio control).
By assigning an expression to a variable name, you can use that expression by its name to compose other objects.

### variables

Variable names must start with a non-digit character and can have zero or more characters in [a-z A-Z _ 0-9].
An assignment "=" is used to create a variable.
To delete a variable, assign it to the special value "nil".

### comment

Use "//" to add comment, either on a new line or and the end of an expression.

## Creation functions

- <a href="#chord">chord</a>
- <a href="#chordsequence">chordsequence</a>
- <a href="#midi">midi</a>
- <a href="#note">note</a>
- <a href="#progression">progression</a>
- <a href="#scale">scale</a>
- <a href="#sequence">sequence</a>

## Composition functions

- <a href="#at">at</a>
- <a href="#duration">duration</a>
- <a href="#dynamic">dynamic</a>
- <a href="#dynamicmap">dynamicmap</a>
- <a href="#export">export</a>
- <a href="#fraction">fraction</a>
- <a href="#fractionmap">fractionmap</a>
- <a href="#group">group</a>
- <a href="#import">import</a>
- <a href="#interval">interval</a>
- <a href="#iterator">iterator</a>
- <a href="#join">join</a>
- <a href="#joinmap">joinmap</a>
- <a href="#listen">listen</a>
- <a href="#merge">merge</a>
- <a href="#midi_send">midi_send</a>
- <a href="#next">next</a>
- <a href="#notemap">notemap</a>
- <a href="#octave">octave</a>
- <a href="#octavemap">octavemap</a>
- <a href="#onbar">onbar</a>
- <a href="#pitch">pitch</a>
- <a href="#print">print</a>
- <a href="#random">random</a>
- <a href="#repeat">repeat</a>
- <a href="#replace">replace</a>
- <a href="#resequence">resequence</a>
- <a href="#reverse">reverse</a>
- <a href="#stretch">stretch</a>
- <a href="#track">track</a>
- <a href="#transpose">transpose (pitch)</a>
- <a href="#transposemap">transposemap (pitchmap)</a>
- <a href="#undynamic">undynamic</a>
- <a href="#ungroup">ungroup</a>
- <a href="#velocitymap">velocitymap</a>

## Audio control functions

- <a href="#biab">biab</a>
- <a href="#bpm">bpm</a>
- <a href="#channel">channel</a>
- <a href="#device">device</a>
- <a href="#key">key</a>
- <a href="#knob">knob</a>
- <a href="#loop">loop</a>
- <a href="#multitrack">multitrack</a>
- <a href="#onkey">onkey</a>
- <a href="#play">play</a>
- <a href="#record">record</a>
- <a href="#set">set</a>
- <a href="#stop">stop</a>
- <a href="#sync">sync</a>


### at
<a name="at"></a>
Create an index getter (1-based) to select a musical object.

> at(index,object)
	
```javascript
at(1,scale('e/m')) // => E
```

### biab
<a name="biab"></a>
Set the Beats in a Bar; default is 4.

> biab(beats-in-a-bar)
	
```javascript
biab(4)
```

### bpm
<a name="bpm"></a>
Set the Beats Per Minute (BPM) [1..300]; default is 120.

> bpm(beats-per-minute)
	
```javascript
bpm(90)

speedup = iterator(80,100,120,140)

l = loop(bpm(speedup),sequence('c e g'),next(speedup))
```

### channel
<a name="channel"></a>
Select a MIDI channel, must be in [1..16]; must be a top-level operator.

> channel(number,sequenceable)
	
```javascript
channel(2,sequence('c2 e3')) // plays on instrument connected to MIDI channel 2
```

### chord
<a name="chord"></a>
Create a Chord from its string <a href="/docs/reference/notations/#chord">format</a>.

> chord('note')
	
```javascript
chord('c#5/m/1')

chord('g/M/2') // Major G second inversion
```

### chordsequence
<a name="chordsequence"></a>
Create a Chord sequence using this <a href="/docs/reference/notations/#chordsequence">format</a>.

> chordsequence('chords')
	
```javascript
chordsequence('e f') // => (E A_ B) (F A C5)

chordsequence('(c d)') // => (C E G D G_ A)
```

### device
<a name="device"></a>
Select a MIDI device from the available device IDs; must become before channel.

> device(number,sequenceable)
	
```javascript
device(1,channel(2,sequence('c2 e3'))) // plays on connected device 1 through MIDI channel 2
```

### duration
<a name="duration"></a>
Computes the duration of the object using the current BPM.

> duration(object)
	
```javascript
duration(note('c')) // => 375ms
```

### dynamic
<a name="dynamic"></a>
Creates a new modified musical object for which the dynamics of all notes are changed.
	The first parameter controls the emphasis the note, e.g. + (mezzoforte,mf), -- (piano,p) or a velocity [0..127].
	.

> dynamic(emphasis,object)
	
```javascript
dynamic('++',sequence('e f')) // => E++ F++

dynamic(112,note('a')) // => A++++
```

### dynamicmap
<a name="dynamicmap"></a>
Changes the dynamic of notes from a musical object. 1-index-based mapping.

> dynamicmap('mapping',object)
	
```javascript
dynamicmap('1:++,2:--',sequence('e f')) // => E++ F--

dynamicmap('2:o,1:++,2:--,1:++', sequence('a b') // => B A++ B-- A++
```

### export
<a name="export"></a>
Writes a multi-track MIDI file.

> export(filename,sequenceable)
	
```javascript
export('myMelody-v1',myObject)
```

### fraction
<a name="fraction"></a>
Creates a new object for which the fraction of duration of all notes are changed.
The first parameter controls the fraction of the note, e.g. 1 = whole, 2 = half, 4 = quarter, 8 = eight, 16 = sixteenth.
Fraction can also be an exact float value between 0 and 1.
.

> fraction(object,object)
	
```javascript
fraction(8,sequence('e f')) // => 8E 8F , shorten the notes from quarter to eight
```

### fractionmap
<a name="fractionmap"></a>
Create a sequence with notes for which the fractions are changed. 1-based indexing. use space or comma as separator.

> fractionmap('fraction-mapping',object)
	
```javascript
fractionmap('3:. 2:4,1:2',sequence('c e g')) // => .G E 2C
```

### group
<a name="group"></a>
Create a new sequence in which all notes of a musical object are grouped.

> group(sequenceable)
	
```javascript
group(sequence('c d e')) // => (C D E)
```

### import
<a name="import"></a>
Evaluate all the statements from another file.

> import(filename)
	
```javascript
import('drumpatterns.mel')
```

### interval
<a name="interval"></a>
Create an integer repeating interval (from,to,by,method). Default method is 'repeat', Use next() to get a new integer.

> interval(from,to,by)
	
```javascript
int1 = interval(-2,4,1)

lp_cdef = loop(transpose(int1,sequence('c d e f')), next(int1))
```

### iterator
<a name="iterator"></a>
Iterator that has an array of constant values and evaluates to one. Use next() to increase and rotate the value.

> iterator(array-element)
	
```javascript
i = iterator(1,3,5,7,9)

p = transpose(i,note('c'))

lp = loop(p,next(i))

		
```

### join
<a name="join"></a>
Joins one or more musical objects as one.

> join(first,second)
	
```javascript
a = chord('a')

b = sequence('(c e g)')

ab = join(a,b) // => (A D_5 E5) (C E G)
```

### joinmap
<a name="joinmap"></a>
Creates a new join by mapping elements. 1-index-based mapping.

> joinmap('indices',join)
	
```javascript
j = join(note('c'), sequence('d e f'))

jm = joinmap('1 (2 3) 4',j) // => C = D =
```

### key
<a name="key"></a>
Use the key to trigger the play of musical object.

> key('note')
	
```javascript
c2 = key('c2') // C2 key on the default input device and default channel

c2 = key(device(1,note('c2'))) // C2 key on input device 1

c2 = key(device(1,channel(2,note('c2'))) // C2 key on input device 1 and channel 2

c2 = key(channel(3,note('c2')) // C2 key on the default input device and channel 3
```

### knob
<a name="knob"></a>
Use the knob as an integer value for a parameter in any object.

> knob(device-id,midi-number)
	
```javascript
axiom = 1 // device ID for my connected M-Audio Axiom 25

B1 = 20 // MIDI number assigned to this knob on the controller

k = knob(axiom,B1)

transpose(k,scale(1,'E')) // when played, use the current value of knob "k"
```

### listen
<a name="listen"></a>
Listen for note(s) from a device and call a playable function to handle.

> listen(variable-or-device-selector,function)
	
```javascript
rec = note('c') // define a variable "rec" with a initial object ; this is a place holder

fun = play(rec) // define the playable function to call when notes are received ; loop and print are also possible

listen(rec,fun) // start a listener for notes from default input device, store it in "rec" and call "fun"

listen(device(1,rec),fun) // start a listener for notes from input device 1
```

### loop
<a name="loop"></a>
Create a new loop from one or more musical objects.

> loop(object)
	
```javascript
cb = sequence('c d e f g a b')

loop(cb,reverse(cb))
```

### merge
<a name="merge"></a>
Merges multiple sequences into one sequence.

> merge(sequenceable)
	
```javascript
m1 = notemap('..!..!..!', note('c2'))

m2 = notemap('4 7 10', note('d2'))

all = merge(m1,m2) // => = = C2 D2 = C2 D2 = C2 D2 = =
```

### midi
<a name="midi"></a>
Create a Note from MIDI information and is typically used for drum sets.
The first parameter is a fraction {1,2,4,8,16} or a duration in milliseconds or a time.Duration.
Second parameter is the MIDI number and must be one of [0..127].
The third parameter is the velocity (~ loudness) and must be one of [0..127].

> midi(numberOrDuration,number,number)
	
```javascript
midi(500,52,80) // => 500ms E3+

midi(16,36,70) // => 16C2 (kick)
```

### midi_send
<a name="midi_send"></a>
Sends a MIDI message with status, channel(ignore if < 1), 2nd byte and 3rd byte to an output device. Can be used as a musical object.

> midi_send(device-id,status,channel,2nd-byte,3rd-byte
	
```javascript
midi_send(1,0xB0,7,0x7B,0) // to device id 1, control change, all notes off in channel 7

midi_send(1,0xC0,2,1,0) // program change, select program 1 for channel 2

midi_send(2,0xB0,4,0,16) // control change, bank select 16 for channel 4

midi_send(3,0xB0,1,120,0) // control change, all notes off for channel 1
```

### multitrack
<a name="multitrack"></a>
Create a multi-track object from zero or more tracks.

> multitrack(track)
	
```javascript
multitrack(track1,track2,track3) // 3 tracks in one multi-track object
```

### next
<a name="next"></a>
Is used to produce the next value in a generator such as random, iterator and interval.
The function itself does not return the value; use the generator for that.

> 
	
```javascript
i = interval(-4,4,2)

pi = transpose(i,sequence('c d e f g a b')) // current value of "i" is used

lp_pi = loop(pi,next(i)) // "i" will advance to the next value

begin(lp_pi)
```

### note
<a name="note"></a>
Create a Note using this <a href="/docs/reference/notations/#note">format</a>.

> note('letter')
	
```javascript
note('e')

note('2.e#--')
```

### notemap
<a name="notemap"></a>
Creates a mapper of notes by index (1-based) or using dots (.) and bangs (!).

> notemap('space-separated-1-based-indices-or-dots-and-bangs',has-note)
	
```javascript
m1 = notemap('..!..!..!', note('c2'))

m2 = notemap('3 6 9', octave(-1,note('d2')))
```

### octave
<a name="octave"></a>
Change the pitch of notes by steps of 12 semitones for one or more musical objects.

> octave(offset,sequenceable)
	
```javascript
octave(1,sequence('c d')) // => C5 D5
```

### octavemap
<a name="octavemap"></a>
Create a sequence with notes for which the order and the octaves are changed.

> octavemap('int2int',object)
	
```javascript
octavemap('1:-1,2:0,3:1',chord('c')) // => (C3 E G5)
```

### onbar
<a name="onbar"></a>
Puts a musical object on a track to start at a specific bar.

> onbar(bar,object)
	
```javascript
tr = track("solo",2, onbar(1,soloSequence)) // 2 = channel
```

### onkey
<a name="onkey"></a>
Assign a playable to a key.
If this key is pressed the playable will start. 
If pressed again, the play will stop.
Remove the assignment using the value nil for the playable.

> onkey(key,playable-or-evaluatable-or-nil)
	
```javascript
axiom = 1 // device ID for the M-Audio Axiom 25

c2 = key(axiom,'c2')

fun = play(scale(2,'c')) // what to do when a key is pressed (NoteOn)

onkey(c2, fun) // if C2 is pressed on the axiom device that evaluate the function "fun"
```

### pitch
<a name="pitch"></a>
Change the pitch with a delta of semitones.

> transpose(semitones,sequenceable)
	
```javascript
transpose(-1,sequence('c d e'))

p = interval(-4,4,1)

transpose(p,note('c'))
```

### play
<a name="play"></a>
Play all musical objects.

> play(sequenceable)
	
```javascript
play(s1,s2,s3) // play s3 after s2 after s1
```

### print
<a name="print"></a>
Prints an object when evaluated (play,loop).

> 
	
```javascript

```

### progression
<a name="progression"></a>
Create a Chord progression using this <a href="/docs/reference/notations/#chordprogression">format</a>.

> progression('scale','space-separated-roman-chords')
	
```javascript
progression('C','II V I') // => (D F A) (G B D5) (C E G)
```

### random
<a name="random"></a>
Create a random integer generator. Use next() to generate a new integer.

> random(from,to)
	
```javascript
num = random(1,10)

next(num)
```

### record
<a name="record"></a>
Create a recorded sequence of notes from the current MIDI input device.

> record(rec)
	
```javascript
rec = sequence('') // variable to store the recorded sequence

record(rec) // record notes played on the current input device and stop recording after 5 seconds
```

### repeat
<a name="repeat"></a>
Repeat one or more musical objects a number of times.

> repeat(times,sequenceables)
	
```javascript
repeat(4,sequence('c d e'))
```

### replace
<a name="replace"></a>
Replaces all occurrences of one musical object with another object for a given composed musical object.

> replace(target,from,to)
	
```javascript
c = note('c')

d = note('d')

pitchA = transpose(1,c)

pitchD = replace(pitchA, c, d) // c -> d in pitchA
```

### resequence
<a name="resequence"></a>
Creates a modifier of sequence notes by index (1-based).

> resequence('space-separated-1-based-indices',sequenceable)
	
```javascript
s1 = sequence('C D E F G A B')

i1 = resequence('6 5 4 3 2 1',s1) // => B A G F E D

i2 = resequence('(6 5) 4 3 (2 1)',s1) // => (B A) G F (E D)
```

### reverse
<a name="reverse"></a>
Reverse the (groups of) notes in a sequence.

> reverse(sequenceable)
	
```javascript
reverse(chord('a'))
```

### scale
<a name="scale"></a>
Create a Scale using this <a href="/docs/reference/notations/#scale">format</a>.

> scale(octaves,'scale-syntax')
	
```javascript
scale(1,'e/m') // => E F G A B C5 D5
```

### sequence
<a name="sequence"></a>
Create a Sequence using this <a href="/docs/reference/notations/#sequence">format</a>.

> sequence('space-separated-notes')
	
```javascript
sequence('c d e')

sequence('(8c d e)') // => (8C D E)

sequence('c (d e f) a =')
```

### set
<a name="set"></a>
Generic function to change a default setting.

> set(setting-name,setting-value)
	
```javascript
set('midi.in',1) // default MIDI input device is 1

set('midi.in.channel',2,10) // default MIDI channel for device 2 is 10

set('midi.out',3) // default MIDI output device is 3
```

### stop
<a name="stop"></a>
Stop running loop(s) or listener(s). Ignore if it was stopped.

> stop(control)
	
```javascript
l1 = loop(sequence('c e g'))

play(l1)

stop(l1)

stop() // stop all playables
```

### stretch
<a name="stretch"></a>
Stretches the duration of musical object(s) with a factor. If the factor < 1 then duration is shortened.

> stretch(factor,object)
	
```javascript
stretch(2,note('c'))  // 2C

stretch(0.25,sequence('(c e g)'))  // (16C 16E 16G)

stretch(8,note('c'))  // C with length of 8 x 0.25 (quarter) = 2 bars
```

### sync
<a name="sync"></a>
Synchronise playing musical objects. Use play() for serial playing.

> sync(object)
	
```javascript
sync(s1,s2,s3) // play s1,s2 and s3 at the same time

sync(loop1,loop2) // begin loop2 at the next start of loop1
```

### track
<a name="track"></a>
Create a named track for a given MIDI channel with a musical object.

> track('title',midi-channel, onbar(1,object))
	
```javascript
track("lullaby",1,onbar(2, sequence('c d e'))) // => a new track on MIDI channel 1 with sequence starting at bar 2
```

### transpose (pitch)
<a name="transpose"></a>
Change the pitch with a delta of semitones.

> transpose(semitones,sequenceable)
	
```javascript
transpose(-1,sequence('c d e'))

p = interval(-4,4,1)

transpose(p,note('c'))
```

### transposemap (pitchmap)
<a name="transposemap"></a>
Create a sequence with notes for which the order and the pitch are changed. 1-based indexing.

> transposemap('int2int',object)
	
```javascript
transposemap('1:-1,1:0,1:1',note('c')) // => B3 C D
```

### undynamic
<a name="undynamic"></a>
Set the dymamic to normal for all notes in a musical object.

> undynamic(sequenceable)
	
```javascript
undynamic('A+ B++ C-- D-') // =>  A B C D
```

### ungroup
<a name="ungroup"></a>
Undo any grouping of notes from one or more musical objects.

> ungroup(sequenceable)
	
```javascript
ungroup(chord('e')) // => E G B

ungroup(sequence('(c d)'),note('e')) // => C D E
```

### velocitymap
<a name="velocitymap"></a>
Create a sequence with notes for which the order and the velocities are changed. Velocity 0 means no change.

> velocitymap('int2int',object)
	
```javascript
velocitymap('1:30,2:0,3:60',chord('c')) // => (C3--- E G5+)
```



##### generated by dsl-md.go on May 22, 2021
