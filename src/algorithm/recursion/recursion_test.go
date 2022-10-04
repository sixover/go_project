package recursion

import (
	"fmt"
	"testing"
)

func TestGetIntNumMax(t *testing.T) {
	arr := []int{22, 123, 1, 23, 1, 4, 235, 1, 43, 6, 37, 145, 7, 345}
	max := getIntNumMax(arr, 0, len(arr)-1)
	fmt.Println(max)
}
