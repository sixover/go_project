package xor_test

import (
	"fmt"
	"testing"
)

func TestFindTwoJi(t *testing.T) {
	arr := []int{1, 1, 2, 2, 2, 3, 3, 4, 4, 4, 4, 5, 6, 6}
	xor1 := 0
	for _, res := range arr {
		xor1 = xor1 ^ res
	}
	importantBit := xor1 & ((^xor1) + 1)
	xor2 := 0
	for _, res := range arr {
		if res&importantBit != 0 {
			xor2 = xor2 ^ res
		}
	}
	fmt.Println(xor2, " ", xor2^xor1)
}
