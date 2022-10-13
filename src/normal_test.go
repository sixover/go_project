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
	res := &ListNode{}
	fmt.Println(res)
}
