package stack_and_queue

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	dl := new(doubleLink)
	for i := 0; i < 10; i++ {
		addOneNodeFromHead(dl, i)
	}
	for i := 0; i < 10; i++ {
		j := getOneNodeFromHead(dl)
		fmt.Println(j)
	}
	for i := 0; i < 10; i++ {
		addOneNodeFromTail(dl, i)
	}
	for i := 0; i < 10; i++ {
		j := getOneNodeFromTail(dl)
		fmt.Println(j)
	}

	stack := getStackInit()
	for i := 0; i < 10; i++ {
		stack.push(i)
	}
	fmt.Println(stack.arr)
	for i := 0; i < 10; i++ {
		j := stack.pop()
		fmt.Println(j)
	}
}
