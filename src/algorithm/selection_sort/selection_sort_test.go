package selection_sort

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	a := []int{7, 8, 1, 888, 54, 654, 3, 213, 1, 65, 487}
	SelectionSort(a)
	fmt.Println(a)
}

func SelectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min != i {
			swap(arr, i, min)
		}
	}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
