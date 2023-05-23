package heap_sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestHeapStruct(t *testing.T) {
	rand.Seed(time.Now().Unix())
	heap := new(heapSmall)
	for i := 0; i < 10; i++ {
		heap.push(rand.Intn(100))
	}
	fmt.Println(heap.heap)
	for i := 0; i < 10; i++ {
		res := heap.pop()
		fmt.Println(res)
	}
}

func TestHeapSort(t *testing.T) {
	rand.Seed(time.Now().Unix())
	a := []int{}
	for i := 0; i < 10; i++ {
		a = append(a, rand.Intn(100))
	}
	fmt.Println(a)

	//建堆
	//第一种方式，每次push一个切片中的元素进来
	for i := 0; i < len(a); i++ {
		heapIn(a, i)
	}
	//第二种方式，从后往前每次进行调整
	//for i := len(a) - 1; i >= 0; i-- {
	//	adjs(a, i, len(a))
	//}

	fmt.Println(a)

	temp := len(a)
	for i := 0; i < temp; i++ {
		a[0], a[len(a)-1] = a[len(a)-1], a[0]
		a = a[:len(a)-1]
		adjs(a, 0, len(a))
		fmt.Println(a)
	}
	//建堆之后，每次取出堆顶，然后与数组最后交换位置，再次调整
	//for i := len(a) - 1; i > 0; i-- {
	//	swap(a, 0, i)
	//	heapAdjust(a, 0, i)
	//}
	//fmt.Println(a)
}

func TestMostOkSort(t *testing.T) {
	a := []int{3, 4, 1, 2, 5}
	sortedArrDistanceLessK(a, 2)
	fmt.Println(a)
}
