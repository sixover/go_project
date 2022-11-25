package radix_sort

import (
	"fmt"
	"sort"
	"testing"
)

func TestRadixSort(t *testing.T) {
	a := []int{1, 3, 53, 2, 23, 46, 2, 346, 347, 22}
	sort.Ints(a)
	fmt.Println(a)
	b := []int{1, 3, 53, 2, 23, 46, 2, 346, 347, 22}
	radixSort(b, 0, len(b)-1)
	fmt.Println(b)
	c := []int{1, 3, 53, 2, 23, 46, 2, 346, 347, 22}
	radixSort2(c, 0, len(c)-1)
	fmt.Println(c)
}
