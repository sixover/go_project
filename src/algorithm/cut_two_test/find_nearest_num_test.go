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

}
