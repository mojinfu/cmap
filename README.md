# cmap 

## support feather

1. support O(1) length method,  like the issue decsribed here https://github.com/golang/go/issues/20680

		新增特性，O(1)复杂度的长度方法

2. reduce contention between Map operations with muti hash lock (doing),  like the issue decsribed here https://github.com/golang/go/issues/21035

		目前sync map 对不同的key操作，对应的都是同一个锁，可以通过多锁的方式，减少竞争。（实现中）

In Go 1.9, `sync.Map` was introduced, but  `sync.Map`  dont support length method , 

when you need get the length of your map

获取map长度示例对比

## sync.map usage 
```go

import (
	"sync"
)

length := 0

myMap.Range(func(_, _ interface{}) bool {
	length++
	return true
})
```
it will lock your map, and take O(n) times

此操作会触发锁，复杂度O(n)

## cmap usage 

```go
import (
	"github.com/mojinfu/cmap"
)

length := myCMap.Length()
```
cmap will **not** lock your cmap, and take O(1) times

此操作**不**会触发锁，复杂度O(1)

## benchmark

100 times Store(i, i) and Delete(i) in env  goos: darwin ; goarch: amd64

 | package | ns/op| B/op|allocs/op|
| :------:| :------: | :------: | :------: |
| sync.Map| 21230 ns/op|	5600 B/op|	499 allocs/op|
| cmap.Map |24243 ns/op|	5600 B/op	|499 allocs/op|
 ----------

 - it means each Store or Delete action will take another 15ns
 - cmap.Map中使用的原子计数器虽然线程安全，是通过底层硬件的支持作为保障的，这使得每次新增Key 删除Key，相对于sync.Map 都将有15ns的额外耗时

  | map length | get length in sync.Map time-consuming |store in cmap.Map Extra time consuming |
| :------:| :------: | :------: |
| map长度 |  使用sync.Map通过range获取长度耗时 |使用CMap存删造成的理论额外耗时 |
| 1 |  44.4 ns |15 ns |
| 10 |  168 ns |150 ns |
| 100 |  1527 ns |1500 ns |
| 1000 |  22647 ns |15000 ns |
| 10000 |  0.32 ms |0.15 ms |
| 100000 |  8.01 ms |1.5 ms |
 ----------

  - it means when programme needs the length of the map(even for one time), use cmap package take the place of sync.map is better.

  - 如果程序中需要获取Map长度，务必使用cmap来减少性能损耗。对长度的获取越频繁，使用cmap的必要性就越大。
