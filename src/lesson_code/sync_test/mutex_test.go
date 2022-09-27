package sync_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	count := 0
	var mut sync.Mutex
	for i := 0; i < 10000; i++ {
		go func() {
			defer mut.Unlock()
			mut.Lock()
			count++
		}()
	}
	time.Sleep(time.Second * 1)
	fmt.Println(count)
}
