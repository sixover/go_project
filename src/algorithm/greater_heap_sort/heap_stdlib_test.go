package greater_heap_sort

import (
	"container/heap"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

type student struct {
	name string
	age  int
}

type stuHeap []student

func (sh *stuHeap) Len() int {
	return len(*sh)
}
func (sh *stuHeap) Less(i, j int) bool {
	if (*sh)[i].age != (*sh)[j].age {
		return (*sh)[i].age < (*sh)[j].age
	} else {
		return (*sh)[i].name < (*sh)[j].name
	}

}
func (sh *stuHeap) Swap(i, j int) {
	(*sh)[i], (*sh)[j] = (*sh)[j], (*sh)[i]
}
func (sh *stuHeap) Push(x interface{}) {
	*sh = append(*sh, x.(student))
}
func (sh *stuHeap) Pop() interface{} {
	ret := (*sh)[sh.Len()-1]
	*sh = (*sh)[:sh.Len()-1]
	return ret
}

func TestHeapStdlib(t *testing.T) {
	h := new(stuHeap)
	for i := 0; i < 10; i++ {
		t := rand.Intn(20)
		s := student{
			name: "stu_" + strconv.Itoa(i),
			age:  t,
		}
		h.Push(s)
	}

	fmt.Println("---")
	fmt.Println(h)

	heap.Init(h)
	fmt.Println(h)

	fmt.Println(heap.Remove(h, 2))
	fmt.Println(h)

	(*h)[1].age = 15
	fmt.Println(h)
	heap.Fix(h, 1)
	fmt.Println(h)

	for i := h.Len(); i > 0; i-- {
		fmt.Println(heap.Pop(h))
	}
}
