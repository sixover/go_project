package channel_test

import (
	"fmt"
	"sync"
	"testing"
)

func producer(ch chan<- int, wg *sync.WaitGroup) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
	wg.Done()
}
func customer(ch <-chan int, wg *sync.WaitGroup) {
	for {
		if data, ok := <-ch; ok {
			fmt.Println(data)
		} else {
			break
		}
	}
	wg.Done()
}

func TestChannelClose(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	go producer(ch, &wg)
	wg.Add(1)
	go customer(ch, &wg)
	wg.Add(1)
	go customer(ch, &wg)
	wg.Wait()
}
