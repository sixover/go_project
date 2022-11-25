package count_sort

import (
	"fmt"
	"sort"
	"testing"
)

func TestCountSort(t *testing.T) {
	a := []int{1, 3, 53, 2, 23, 46, 2, 346, 347, 22}
	sort.Ints(a)
	fmt.Println(a)
	b := []int{1, 3, 53, 2, 23, 46, 2, 346, 347, 22}
	countSort(b)
	fmt.Println(b)
}
