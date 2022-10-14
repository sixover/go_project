package quick_sort

import (
	"fmt"
	"testing"
)

func TestGuoQi(t *testing.T) {
	arr := []int{1, 234, 1, 235, 3, 452, 34}
	fmt.Println(flagSimpleSort(arr, 7))
	arr = []int{1, 234, 7, 1, 235, 7, 3, 452, 34}
	fmt.Println(flagComplexSort(arr, 7))
	arr = []int{1, 4, 2, 3, 3}
	quickSort(arr)
	fmt.Println(arr)
	arr = []int{1, 234, 1, 235, 3, 452, 34}
	quickSort(arr)
	fmt.Println(arr)
	arr = []int{1, 234, 7, 1, 235, 7, 3, 452, 34}
	quickSort(arr)
	fmt.Println(arr)
}
