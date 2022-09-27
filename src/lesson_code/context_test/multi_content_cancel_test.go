package context_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func ContentCancel(ch chan interface{}) {
	close(ch)
}

func IsTestCancel(ch chan interface{}) bool {
	select {
	case <-ch:
		return true
	default:
		return false
	}
}
func TestMultiContentCancel(t *testing.T) {
	ch := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int, ch chan interface{}, wg *sync.WaitGroup) {
			for {
				if IsTestCancel(ch) {
					break
				}
				time.Sleep(time.Millisecond * 100)
			}
			fmt.Println(i, " is Done!")
			wg.Done()
		}(i, ch, &wg)
	}
	ContentCancel(ch)
	wg.Wait()
}
