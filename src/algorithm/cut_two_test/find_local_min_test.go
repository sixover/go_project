package cut_two

import (
	"fmt"
	"testing"
)

func TestFindLocalMin(t *testing.T) {
	a := []int{8, 8, 1, 888, 54, 654, 3, 213, 1, 65, 487}
	exist := FindLocalMin(a)
	fmt.Println(exist)
}
func FindLocalMin(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	if len(arr) == 1 {
		return 0
	}
	if arr[0] < arr[1] {
		return 0
	}
	if arr[len(arr)-1] < arr[len(arr)-2] {
		return len(arr) - 1
	}
	l := 1
	r := len(arr) - 2
	for {
		mid := l + (r-l)/2
		if arr[mid] < arr[mid-1] && arr[mid] < arr[mid+1] {
			return mid
		}
		if arr[mid] > arr[mid-1] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
}
