
```javascript
it = iterator(6,1,4,-1) // delta semitones
pm = pitchmap('1:0,1:7,1:12,1:15,1:19,1:24',note('c')) // six notes
sm = fraction(8,resequence('1 4 2 4 3 6 5 4',pm)) // note sequence of eights
lp_sq = loop(pitch(it,sm),next(it)) // loop the sequence and change pitch on every iteration
```
