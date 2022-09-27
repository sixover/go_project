package mapUsing

import "testing"

func TestMapFactory(t *testing.T) {
	fac := make(map[int]func(int) int)
	fac[1] = func(op int) int {
		return op
	}
	fac[2] = func(op int) int {
		return op * op
	}
	fac[3] = func(op int) int {
		return op * op * op
	}
	t.Log(fac[1](2), fac[2](2), fac[3](2))
}
