package sync

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type singleTon struct {
}

var single *singleTon
var once = sync.Once{}

func runOnce() *singleTon {
	once.Do(
		func() {
			single = new(singleTon)
			fmt.Println("this is first creat singleTon")
		})
	return single
}
func TestSyncOnce(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			sin := runOnce()
			fmt.Printf("%x\n", unsafe.Pointer(sin))
			wg.Done()
		}()
	}
	wg.Wait()
}

type aaaa struct {
	x int
}

func (a aaaa) testfunc() {
	fmt.Println("aaaa print!")
}
func TestSyncOnceType(t *testing.T) {
	var once sync.Once
	var a aaaa
	a.testfunc()
	fmt.Printf("%T\n", once)
}
