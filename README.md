# cmap

In Go 1.9, `sync.Map` was introduced.

but  `sync.Map`  dont support length method , as the issue decsribed https://github.com/golang/go/issues/20680

when you need get the length of your map

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

## cmp usage

```go
import (
	"github.com/mojinfu/cmap"
)

length := myCMap.Length()
```
cmap will **not** lock your cmap, and take O(1) times


## get package
```bash
go get "github.com/mojinfu/cmap"
```

The package is now imported under the "cmap" namespace.

