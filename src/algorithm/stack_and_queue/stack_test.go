package stack_and_queue

import (
	"fmt"
	"testing"
)

func TestStackDoubleChain(t *testing.T) {
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
}
func TestNumSlice(t *testing.T) {
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
func TestMinStack(t *testing.T) {
	stack := getMinStackInit()
	for i := 0; i < 10; i++ {
		stack.push(i)
	}
	fmt.Println(stack.min)
	fmt.Println(stack.data)
	for i := 0; i < 10; i++ {
		s := stack.getMinValue()
		j := stack.pop()
		fmt.Println(j)
		fmt.Println(s)
	}
}

func TestStackUsingByQueue(t *testing.T) {
	stack := getStackByQueue()
	for i := 0; i < 10; i++ {
		stack.push(i)
	}
	for i := 0; i < 10; i++ {
		j := stack.pop()
		fmt.Println(j)
	}
}
