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
func TestRightSmallNum(t *testing.T) {
	a1 := []int{2, 1, 3, 4}
	a2 := []int{2, 1, 3, 4}
	res1 := findNumOfRightSmaller1(a1, 0, len(a1)-1)
	res2 := findNumOfRightSmaller2(a2, 0, len(a2)-1)
	fmt.Println(res1)
	fmt.Println(res2)
	b1 := []int{7, 8, 1, 888, 54, 654, 3, 213, 1, 65, 487}
	res1 = findNumOfRightSmaller1(b1, 6, len(b1)-1)
	b2 := []int{7, 8, 1, 888, 54, 654, 3, 213, 1, 65, 487}
	res2 = findNumOfRightSmaller2(b2, 6, len(b2)-1)
	fmt.Println(res1)
	fmt.Println(res2)
}
func TestRightSmallDoubleNum(t *testing.T) {
	a := []int{7, 1, 3, 4}
	res := findNumOfRightSmallerDoubleNum(a, 0, len(a)-1)
	fmt.Println(res)
	b := []int{7, 8, 1, 888, 54, 654, 3, 213, 1, 65, 487}
	res = findNumOfRightSmallerDoubleNum(b, 0, 4)
	fmt.Println(res)
}
func TestCountOfRangeSum(t *testing.T) {
	nums := []int{0, 0}
	fmt.Println(countRangeSum(nums, 0, 0))
}
