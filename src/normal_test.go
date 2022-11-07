package main

import (
	"fmt"
	"testing"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestNormalSomething(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(time.Now().Add(time.Duration(5) * time.Second))
	}
}
