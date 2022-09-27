package context_test

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func ContestCancel(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}
func TestContest(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int, ctx context.Context, wg *sync.WaitGroup) {
			for {
				if ContestCancel(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 50)
			}
			fmt.Println(i, " is Done!")
			wg.Done()
		}(i, ctx, &wg)
	}
	cancel()
	wg.Wait()
}
