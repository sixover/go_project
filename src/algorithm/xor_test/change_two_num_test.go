package xor_test

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	a := 1
	b := 2
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println(a, " ", b)
}
