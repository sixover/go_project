package select_test

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func PrintA(ch chan string) {
	for i := 0; i < 50; i++ {
		ch <- "wx is learning"
	}
}

func PrintB(ch chan string) {
	for i := 0; i < 50; i++ {
		ch <- "wx is touch fish"
	}
}

func TestSelect(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			t.Log(p)
		}
	}()
	ch1 := make(chan string)
	ch2 := make(chan string)
	go PrintA(ch1)
	go PrintB(ch2)
	for {
		select {
		case ret1 := <-ch1:
			fmt.Println(ret1)
		case ret2 := <-ch2:
			fmt.Println(ret2)
		case ret3 := <-time.After(time.Second):
			t.Log(ret3)
			panic(errors.New("run time out Error"))
		}
	}
}
