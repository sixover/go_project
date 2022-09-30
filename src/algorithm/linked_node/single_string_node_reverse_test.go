package linked_node_test

import (
	"fmt"
	"testing"
)

type Node struct {
	value int
	next  *Node
}

func TestSingleNodeReverse(t *testing.T) {
	head := &Node{
		value: 1,
		next:  nil,
	}
	maxNum := 5
	for i := maxNum; i > 1; i-- {
		node := new(Node)
		node.value = i
		if head.next == nil {
			head.next = node
			continue
		}
		node.next = head.next
		head.next = node
	}
	tempPrint := head
	fmt.Println("this is origin double chain!")
	for tempPrint != nil {
		fmt.Println(tempPrint.value)
		tempPrint = tempPrint.next
	}
	res := SingleNodeReverse(head)
	fmt.Println("this is reverse double chain!")
	for res != nil {
		fmt.Println(res.value)
		res = res.next
	}
}

func SingleNodeReverse(head *Node) *Node {
	var pre *Node
	var next *Node

	for head != nil {
		next = head.next
		head.next = pre
		pre = head
		head = next
	}
	return pre
}
