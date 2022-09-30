package single_string_node_reverse

import (
	"fmt"
	"testing"
)

type doubleLinkedNode struct {
	value int
	pre   *doubleLinkedNode
	next  *doubleLinkedNode
}

func TestDoubleReverse(t *testing.T) {
	head := &doubleLinkedNode{
		value: 5,
		next:  nil,
		pre:   nil,
	}
	for i := 1; i < 5; i++ {
		node := new(doubleLinkedNode)
		node.value = i
		if head.next == nil {
			head.next = node
			node.pre = head
			continue
		}
		node.next = head.next
		node.pre = head
		head.next = node
		node.next.pre = node
	}
	tempPrint := head
	fmt.Println("this is origin double chain!")
	for tempPrint != nil {
		fmt.Println(tempPrint.value)
		tempPrint = tempPrint.next
	}
	res := DoubleLinkedNodeReverse(head)
	fmt.Println("this is reverse double chain!")
	for res != nil {
		fmt.Println(res.value)
		res = res.next
	}
}

func DoubleLinkedNodeReverse(head *doubleLinkedNode) *doubleLinkedNode {
	var tpre *doubleLinkedNode
	var tnext *doubleLinkedNode
	for head != nil {
		tnext = head.next
		head.next = tpre
		head.pre = tnext
		tpre = head
		head = tnext
	}
	return tpre
}
