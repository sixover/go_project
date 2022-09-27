package channel_test

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(time.Millisecond * 50)
	return fmt.Sprintf("the res from %d\n", id)
}
func firstResponse() string {
	ch := make(chan string)
	for i := 0; i < 10; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	str := <-ch
	close(ch)
	return str
}

func TestClose(t *testing.T) {
	t.Log("before: ", runtime.NumGoroutine())
	t.Log(firstResponse())
	time.Sleep(time.Second)
	t.Log("after: ", runtime.NumGoroutine())
}
