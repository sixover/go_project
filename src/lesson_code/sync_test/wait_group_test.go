package sync

import (
	"fmt"
	"sync"
	"testing"
)

func TestWaitGroup(t *testing.T) {
	count := 0
	var mut sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer mut.Unlock()
			defer wg.Done()
			mut.Lock()
			count++
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
