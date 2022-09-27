package goroutine_test

import (
	"fmt"
	"testing"
	"time"
)

/*
这个函数的主函数没有进行等待，所以创建的协程在未全部执行完的情况下程序退出了
且由于创建的所有的协程都共享了变量i，所以当主函数for循环改变了i的值时，创建的协程打印i的值也会发生变化
*/
func TestGoroutine1(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
}

//主函数等待
func TestGoroutine2(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(1 * time.Second)
}

//这里创建的协程函数是使用Go语言传值的方法将i赋值给a，也就是说他们不再共享同一个变量，所以可以正常打印
func TestGoroutine3(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(a int) {
			fmt.Println(a)
		}(i)
	}
	time.Sleep(1 * time.Second)
}
