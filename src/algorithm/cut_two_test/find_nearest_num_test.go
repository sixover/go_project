package cut_two

import (
	"fmt"
	"testing"
)

func TestNearestFind(t *testing.T) {
	a := []int{1, 1, 1, 2, 2, 2, 2, 2, 3, 3, 4, 4, 5}
	exist := NearestFinder(a, 2)
	fmt.Println(exist)
}

func NearestFinder(arr []int, num int) int {
	if len(arr) == 0 {
		return -1
	}
	if arr[0] > num {
		return -1
	}
	if arr[len(arr)-1] < num {
		return -1
	}
	l := 0
	r := len(arr) - 1
	index := -1
	for {
		mid := l + (r-l)/2
		if arr[mid] >= num {
			index = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
		if r < l {
			break
		}
	}
	return index
}
