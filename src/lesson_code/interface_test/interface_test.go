package interface_test

import (
	"fmt"
	"testing"
)

type Test interface {
	PrintString() string
}

type SimpleStructForTest struct {
}

func (s *SimpleStructForTest) PrintString() string {
	return fmt.Sprintf("this is interface of struct print!")
}

func UsingInterface(t Test) {
	fmt.Println(t.PrintString())
}

func TestInterface(t *testing.T) {
	var interfa Test
	interfa = new(SimpleStructForTest)
	t.Log(interfa.PrintString())

	S := new(SimpleStructForTest)
	UsingInterface(S)
}
