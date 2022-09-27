package bubble_sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	a := []int{7, 8, 1, 888, 54, 654, 3, 213, 1, 65, 487}
	BubbleSort(a)
	fmt.Println(a)
}
func BubbleSort(arr []int) {
	for i := len(arr); i > 0; i-- {
		for j := 1; j < i; j++ {
			if arr[j-1] > arr[j] {
				swap(arr, j-1, j)
			}
		}
	}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
