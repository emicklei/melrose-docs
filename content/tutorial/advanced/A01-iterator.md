---
title: A01 Iterator
menu:
  tutorial:
    parent: "advanced"
---

## iterator

An iterator is used to as a variable that takes the values, one by one, from a given list of numbers or strings.

```
offsets = iterator(1,2,3,4)
ceg = sequence('C E G')
loop(pitch(offsets,ceg),next(offsets))
```

The `next(offsets)` expression is needed to advance the iterator.