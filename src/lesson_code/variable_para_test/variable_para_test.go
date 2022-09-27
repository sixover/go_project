package variable_para_test

import (
	"fmt"
	"testing"
)

func PrintParaType(para ...int) string {
	return fmt.Sprintf("the type is : %T", para)
}

func TestParaType(t *testing.T) {
	t.Log(PrintParaType(1, 2, 3, 4))
	ints := []int{1, 2, 3, 4}
	t.Log(PrintParaType(ints...))
}
func f(i []int) {

}
func g(i ...int) {

}
func TestDifferent(t *testing.T) {
	t.Log(fmt.Sprintf("%T", f))
	t.Log(fmt.Sprintf("%T", g))
}
