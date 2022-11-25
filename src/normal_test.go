package main

import (
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestNormalSomething(t *testing.T) {
	count := 0
	maxNum := 101
	for maxNum != 0 {
		maxNum = maxNum / 10
		count++
	}
	t.Log(count)
}
