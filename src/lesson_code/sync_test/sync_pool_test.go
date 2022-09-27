package sync_test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestPool(t *testing.T) {
	pool := &sync.Pool{New: func() interface{} {
		fmt.Println("Create a new object!")
		return 0
	}}
	v := pool.Get().(int)
	t.Log(v)
	pool.Put(1)
	v = pool.Get().(int)
	t.Log(v)
	pool.Put(v)
	runtime.GC()
	v = pool.Get().(int)
	t.Log(v)
}
