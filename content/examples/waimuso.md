---
title: "Waimuso"
date: 2020-04-27
draft: false
---
	  Logic Pro X


 Multiple loops with different instruments, each on another channel.

```javascript
bpm(120)

s1 = sequence('1c3--- 1e3--- 1a2--- 1d3---')
l1 = loop(channel(2,s1)) // end(l1)   --- black whisper


l2 = loop(channel(4,s1)) // end(l2)   --- subby bass


s2 = sequence('= 8a3 8b3 2c-  = 8c 8d 2e-  = 8d 8e 2f-   = 8c 8b3 2a3-') 
l3 = loop(channel(1,s2)) // end(l3)   --- grand piano + pad
 

s3 = repeat(4,sequence('= g#2+'))
l4 = loop(channel(3,s3)) // end(l4)   --- SoCal
```
