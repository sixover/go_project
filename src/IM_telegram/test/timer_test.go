package test

import (
	"fmt"
	"testing"
	"time"
)

func Timer(delay, tick time.Duration, fun func(string) bool, param string) {
	go func() {
		if fun == nil {
			return
		}
		t := time.NewTimer(delay)
		fmt.Println(time.Now())
		for {
			select {
			case <-t.C:
				fmt.Println(time.Now())
				if fun(param) == false {
					return
				}
				t.Reset(tick)
			}
		}
	}()
}

func TestTimer(t *testing.T) {
	//timer := time.NewTimer(3 * time.Second)
	//fmt.Println("当前时间为:", time.Now())
	//
	//time := <-timer.C //从定时器拿数据
	//fmt.Println("当前时间为:", time)

	//var wg sync.WaitGroup
	//timer := time.NewTimer(100 * time.Second)
	//wg.Add(1)
	//go func() {
	//	fmt.Println("当前时间为:", time.Now())
	//	t := <-timer.C
	//	fmt.Println("当前时间为:", t)
	//	wg.Done()
	//}()
	////重置定时器时间
	//timer.Reset(3 * time.Second)
	//wg.Wait()

	Timer(time.Duration(3)*time.Second, time.Duration(1)*time.Second, func(s string) bool {
		fmt.Println(s)
		return true
	}, "i will make timer more long time!!!")
	for {
		time.Sleep(time.Second)
	}
}
