package xor_test

import (
	"fmt"
	"testing"
)

func TestRightOne(t *testing.T) {
	a := 6
	res := a & ((^a) + 1)
	fmt.Println(res)
}
