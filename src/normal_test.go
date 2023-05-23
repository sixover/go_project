package main

import (
	"fmt"
	"testing"
)

type Student struct {
	name string
}

type printer interface {
	Pa()
}

func (n *Student) Pa() {
	fmt.Println(n.name)
}

func TestNormalSomething(t *testing.T) {
	var a printer = &Student{"asd"}
	a.Pa()
}
