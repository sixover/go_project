package single_string_node_reverse_test

import (
	"fmt"
	"testing"
)

type Node struct {
	value int
	next  *Node
}

func TestSingleNodeReverse(t *testing.T) {
	node_test := &Node{
		value: 3,
		next: &Node{
			value: 2,
			next: &Node{
				value: 1,
				next:  nil,
			},
		},
	}
	res := SingleNodeReverse(node_test)
	temp := res
	for temp != nil {
		fmt.Println(temp.value)
		temp = temp.next
	}
	fmt.Println(res)
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
