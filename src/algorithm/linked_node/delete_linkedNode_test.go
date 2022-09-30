package linked_node_test

import (
	"fmt"
	"testing"
)

func TestDeleteNode(t *testing.T) {
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
	res := DeleteNode(head, 1)
	fmt.Println("this is reverse double chain!")
	for res != nil {
		fmt.Println(res.value)
		res = res.next
	}
}
func DeleteNode(head *Node, num int) *Node {
	for head != nil {
		if head.value != num {
			break
		}
		head = head.next
	}
	var pre *Node
	var fahead *Node

	if head == nil {
		return head
	}
	fahead = head.next
	pre = head
	for fahead != nil {
		if fahead.value == num {
			pre.next = fahead.next
			fahead = fahead.next
			continue
		}
		fahead = fahead.next
		pre = pre.next
	}

	//另一种写法
	//fahead=head
	//pre=head
	//for fahead!=nil{
	//	if fahead.value==num{
	//		pre.next=fahead.next
	//	}else{
	//		pre=fahead
	//	}
	//	fahead=fahead.next
	//}

	return head
}
