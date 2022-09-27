package xor_test

import (
	"fmt"
	"testing"
)

func TestFindJi(t *testing.T) {
	arr := []int{1, 1, 2, 2, 2, 2, 3, 3, 4, 4, 4, 4, 5, 6, 6}
	xor := 0
	for _, res := range arr {
		xor = xor ^ res
	}
	fmt.Println(xor)
}
