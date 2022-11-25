package main

import (
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestNormalSomething(t *testing.T) {
	a := "aksjhfajwhufiaowuehf"
	b := "绿卡就是的哈佛爱说大话覅欧文"
	for i, res := range []rune(a) {
		t.Log(i, ":  ", string(res))
	}
	t.Log("==============")
	for i, res := range []rune(b) {
		t.Log(i, ":  ", string(res))
	}
	aa := make(map[int]int)
	aa[1] = 1
	aa[2] = 2
	if res, ok := aa[1]; ok {
		t.Log(res)
	}
}
