package channel_test

import (
	"fmt"
	"sync"
	"testing"
)

func Producer(ch chan<- int, pg *sync.WaitGroup) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	pg.Done()
}
func Customer(ch <-chan int, wg *sync.WaitGroup) {
	for {
		if data, ok := <-ch; ok {
			fmt.Println(data)
		} else {
			break
		}
	}
	wg.Done()
}

func TestMultiProducer(t *testing.T) {
	var wg sync.WaitGroup
	var pg sync.WaitGroup
	ch := make(chan int)
	pg.Add(1)
	go Producer(ch, &pg)
	pg.Add(1)
	go Producer(ch, &pg)

	wg.Add(1)
	go Customer(ch, &wg)
	wg.Add(1)
	go Customer(ch, &wg)
	wg.Add(1)
	go Customer(ch, &wg)
	pg.Wait()
	close(ch)
	wg.Wait()
}
