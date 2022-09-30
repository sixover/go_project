package stack_and_queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	dl := new(doubleLink)
	for i := 0; i < 10; i++ {
		addOneNodeFromHead(dl, i)
	}
	for i := 0; i < 10; i++ {
		j := getOneNodeFromTail(dl)
		fmt.Println(j)
	}
}
