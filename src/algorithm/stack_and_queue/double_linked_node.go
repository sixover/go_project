package stack_and_queue

type doubleLinkNode struct {
	value int
	pre   *doubleLinkNode
	next  *doubleLinkNode
}

type doubleLink struct {
	head *doubleLinkNode
	tail *doubleLinkNode
}

func getOneNodeFromHead(dl *doubleLink) int {
	if dl.head == nil {
		return -1
	}
	if dl.tail == dl.head {
		temp := dl.head.value
		dl.tail = nil
		dl.head = nil
		return temp
	}
	tNode := dl.head
	dl.head = dl.head.next
	dl.head.pre = nil
	tNode.next = nil
	return tNode.value
}
func getOneNodeFromTail(dl *doubleLink) int {
	if dl.tail == nil {
		return -1
	}
	if dl.tail == dl.head {
		temp := dl.tail.value
		dl.tail = nil
		dl.head = nil
		return temp
	}
	tNode := dl.tail
	dl.tail = dl.tail.pre
	tNode.pre = nil
	dl.tail.next = nil
	return tNode.value
}
func addOneNodeFromHead(dl *doubleLink, nodeValue int) {
	node := new(doubleLinkNode)
	node.value = nodeValue
	if dl.head == nil {
		dl.head = node
		dl.tail = node
	} else {
		node.next = dl.head
		dl.head.pre = node
		dl.head = node
	}
}
func addOneNodeFromTail(dl *doubleLink, nodeValue int) {
	node := new(doubleLinkNode)
	node.value = nodeValue
	if dl.tail == nil {
		dl.tail = node
		dl.head = node
	} else {
		dl.tail.next = node
		node.pre = dl.tail
		dl.tail = node
	}
}
