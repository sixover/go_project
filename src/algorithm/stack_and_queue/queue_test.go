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

	q := getQueue(10)
	for i := 0; i < 10; i++ {
		push(q, i)
	}
	for i := 0; i < 10; i++ {
		j := pop(q)
		fmt.Println(j)
	}
}
func TestQueueByStack(t *testing.T) {
	q := getQueueByStack()
	for i := 0; i < 10; i++ {
		q.push(i)
	}
	for i := 0; i < 6; i++ {
		j := q.pop()
		fmt.Println(j)
	}
	for i := 0; i < 10; i++ {
		q.push(i)
	}
	for i := 0; i < 14; i++ {
		j := q.pop()
		fmt.Println(j)
	}
}
