package cmap

import (
	"sync"
	"testing"
)

func BenchmarkCMapStore(b *testing.B) {
	b.StopTimer()
	myMap := &Map{}
	b.N = 1000000
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for ii := 0; ii < 100; ii++ {
			myMap.Store(i, i)
			myMap.Delete(i)
		}
	}
}
func BenchmarkSyncMapStore(b *testing.B) {
	b.StopTimer()
	myMap := &sync.Map{}
	b.N = 1000000
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for ii := 0; ii < 100; ii++ {
			myMap.Store(i, i)
			myMap.Delete(i)
		}
	}
}

func BenchmarkSyncMapRange(b *testing.B) {
	b.StopTimer()
	myMap := &sync.Map{}
	for i := 0; i < 1; i++ {
		myMap.Store(i, i)
	}
	b.StartTimer()
	for ii := 0; ii < b.N; ii++ {
		myMap.Range(func(_, _ interface{}) bool {
			return true
		})
	}
}

// 1500000
//len 100000 map range
//BenchmarkSyncMapRange-4   	     200	   8010973 ns/op	       0 B/op	       0 allocs/op

//120000
//len 10000 map range
//BenchmarkSyncMapRange-4   	    5000	    325012 ns/op	       0 B/op	       0 allocs/op

//15000
//len 1000 map range
//BenchmarkSyncMapRange-4   	  100000	     22647 ns/op	       0 B/op	       0 allocs/op

//1500
//len 100 map range
//BenchmarkSyncMapRange-4   	 1000000	      1527 ns/op	       0 B/op	       0 allocs/op

//150
//len 10 map range
//BenchmarkSyncMapRange-4   	10000000	       168 ns/op	       0 B/op	       0 allocs/op

//150
//len 1 map range
//BenchmarkSyncMapRange-4   	30000000	        44.4 ns/op	       0 B/op	       0 allocs/op
