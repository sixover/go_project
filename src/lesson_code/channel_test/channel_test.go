package channel_test

import (
	"fmt"
	"testing"
	"time"
)

func ChannelRetFunc(n int, ch chan []int) {
	list := []int{1, 1}
	for i := 2; i < n; i++ {
		list = append(list, list[i-2]+list[i-1])
	}
	ch <- list
	fmt.Println("the cal is over!!", time.Now())
}
func TestChannel1(t *testing.T) {
	ch := make(chan []int)
	ChannelRetFunc(10, ch)
	t.Log(<-ch)
}
func TestChannel2(t *testing.T) {
	ch := make(chan []int)
	go ChannelRetFunc(10, ch)
	time.Sleep(time.Second * 3)
	t.Log(<-ch, time.Now())
}
func TestChannel3(t *testing.T) {
	ch := make(chan []int, 1)
	go ChannelRetFunc(10, ch)
	time.Sleep(time.Second * 3)
	t.Log(<-ch, time.Now())
}
