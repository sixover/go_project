package cut_two_test

import (
	"fmt"
	"testing"
)

func TestFindNumExist(t *testing.T) {
	a := []int{1, 1, 3, 7, 8, 54, 65, 213, 487, 654, 888}
	exist := FindNumExist(a, 654)
	fmt.Println(exist)
}
func FindNumExist(arr []int, num int) int {
	if len(arr) == 0 {
		return -1
	}
	if num < arr[0] {
		return -1
	}
	if num > arr[len(arr)-1] {
		return -1
	}
	l := 0
	r := len(arr) - 1
	for {
		mid := l + (r-l)/2
		if arr[mid] < num {
			l = mid + 1
		} else if arr[mid] > num {
			r = mid - 1
		} else {
			return mid
		}
		if r < l {
			return -1
		}
	}
}
