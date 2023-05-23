package channel_using

import (
	"fmt"
	"sync"
	"testing"
	"time"
	"unsafe"
)

func sender(ch chan int, wg *sync.WaitGroup) {
	for i := 0; i < 10000; i++ {
		ch <- i
	}
	close(ch)
	wg.Done()
}

func recv(ch chan int, wg *sync.WaitGroup) {
	for {
		v, ok := <-ch
		if ok {
			fmt.Println(v)
		} else {
			break
		}
	}
	wg.Done()
}

func worker(taskCh <-chan int) {
	const N = 5
	// 启动 5 个工作协程
	for i := 0; i < N; i++ {
		go func(id int) {
			for {
				task := <-taskCh
				fmt.Printf("finish task: %d by worker %d\n", task, id)
				time.Sleep(time.Second)
			}
		}(i)
	}
}

var wg sync.WaitGroup

func strToBytes(str string) []byte {
	return *(*[]byte)(unsafe.Pointer(&str))
}
func BytesToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func TestChannel(t *testing.T) {
	a := "hello"
	b := strToBytes(a)
	b[0] = 'H'
}
