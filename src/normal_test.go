package main

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestNormalSomething(t *testing.T) {
	a := []int{1, 2, 3, 4}
	fmt.Println(a)
	b := a[:]
	b[1] = 222
	fmt.Println(a)
}
