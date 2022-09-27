package context_test

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func CrankContext(i int, ctxx context.Context, wg *sync.WaitGroup) {
	for {
		if ContestCancel(ctxx) {
			break
		}
		time.Sleep(time.Millisecond * 50)
	}
	fmt.Println(i, " is Done!")
	wg.Done()
}

func BrankContext(i int, ctx context.Context, wg *sync.WaitGroup) {
	ctxx, _ := context.WithCancel(ctx)
	wg.Add(1)
	go CrankContext(i+10, ctxx, wg)

	for {
		if ContestCancel(ctx) {
			break
		}
		time.Sleep(time.Millisecond * 50)
	}
	fmt.Println(i, " is Done!")
	wg.Done()
}

func TestMultiContest(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	//fmt.Printf("%T\n",ctx)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go BrankContext(i, ctx, &wg)
	}
	cancel()
	wg.Wait()
}
