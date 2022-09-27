package func_test

import (
	"fmt"
	"testing"
	"time"
)

func FuncA(b func(int) int) func(int) int {
	return func(i int) int {
		start := time.Now()
		ret := b(i)
		fmt.Println("its spend time: ", time.Since(start).Seconds())
		return ret
	}
}
func FuncB(i int) int {
	time.Sleep(2 * time.Second)
	return i
}
func TestFunc(t *testing.T) {
	getFuncNop := FuncA(FuncB)
	fmt.Printf("%T\n", getFuncNop)
}
