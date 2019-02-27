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
