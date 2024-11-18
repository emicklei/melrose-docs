---
title: "Frere_jacques"
date: 2020-11-16
draft: false
---
 Uses `track` to play sequences together, starting at given bars.

```javascript
f1 = sequence('C D E C')
f2 = sequence('E F 2G')
f3 = sequence('8G 8A 8G 8F E C')
f4 = sequence('2C 2G3 2C 2=') 

v1 = join(f1,f1,f2,f2,f3,f3,f4) 

track('frere',1,
    onbar(1,v1),
    onbar(3,v1),
    onbar(5,v1))
```
