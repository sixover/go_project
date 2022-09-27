package insert_sort

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	a := []int{7, 8, 1, 888, 54, 654, 3, 213, 1, 65, 487}
	InsertSortSecond(a)
	fmt.Println(a)
}
func InsertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		temp := i
		for j := i - 1; j > -1; j-- {
			if arr[temp] < arr[j] {
				swap(arr, temp, j)
				temp = j
			} else {
				break
			}
		}
	}
}
func InsertSortSecond(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i - 1; j > -1; j-- {
			if arr[j+1] < arr[j] {
				swap(arr, j+1, j)
			} else {
				break
			}
		}
	}
}
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
