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

## cmp usage 

```go
import (
	"github.com/mojinfu/cmap"
)

length := myCMap.Length()
```
cmap will **not** lock your cmap, and take O(1) times

此操作**不**会触发锁，复杂度O(1)

## get package
```bash
go get "github.com/mojinfu/cmap"
```

The package is now imported under the "cmap" namespace.

after fixed logic bugs，the result of len method in 'cmap' package is working perfectly
then i found
benchmark : 100 times Store(i, i) and Delete(i) in env  goos: darwin ; goarch: amd64

 | package | ns/op| B/op|allocs/op|
| :------:| :------: | :------: | :------: |
| sync.Map| 21230 ns/op|	5600 B/op|	499 allocs/op|
| CMap |24243 ns/op|	5600 B/op	|499 allocs/op|
 ----------

 it means each Store or Delete action will take another 15ns, I think cmap will be useful in some scene.