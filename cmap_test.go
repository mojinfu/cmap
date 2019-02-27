package cmap

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type Animal struct {
	name string
}

func TestLength(t *testing.T) {
	myMap := &Map{}
	for i := 0; i < 1000000; i++ {
		go myMap.LoadOrStore(rand.Intn(1000000), rand.Intn(1000000))
		go myMap.Store(i, i)
		go myMap.Load(i)
		go myMap.Delete(rand.Intn(1000000))
	}
	time.Sleep(time.Second * 30)
	fmt.Println("O(1) cmp length:", myMap.Length())
	length := 0
	myMap.Range(func(_, _ interface{}) bool {
		length++
		return true
	})
	fmt.Println("O(n) real length:", length)
}
