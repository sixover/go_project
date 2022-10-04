package merge_sort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	a := []int{7, 8, 1, 888, 54, 654, 3, 213, 1, 65, 487}
	b := mergeProcess1(a, 0, len(a)-1)
	fmt.Println(b)
	c := []int{7, 8, 1, 888, 54, 654, 3, 213, 1, 65, 487}
	mergeProcess2(c, 0, len(c)-1)
	fmt.Println(c)
	d := []int{7, 8, 1, 888, 54, 654, 3, 213, 1, 65, 487}
	mergeNonProcess(d)
	fmt.Println(d)
}

func TestSmallSum(t *testing.T) {
	a := []int{2, 1, 3, 4}
	res := small_sum(a, 0, len(a)-1)
	fmt.Println(res)
}
func TestRightBiggerNum(t *testing.T) {
	a := []int{2, 1, 3, 4}
	res := findNumOfRightBigger(a, 0, len(a)-1)
	fmt.Println(res)
	b := []int{7, 8, 1, 888, 54, 654, 3, 213, 1, 65, 487}
	res = findNumOfRightBigger(b, 0, 4)
	fmt.Println(res)
}
