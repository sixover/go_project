package fibo_test

import (
	"math"
	"testing"
)

func TestFibonaci(t *testing.T) {
	var (
		a int = 1
		b     = 1
	)
	for i := 0; i < 5; i++ {
		t.Log(a, b)
		temp := a
		a = b
		b = temp + a
	}
}

func TestMultiTestingRun(t *testing.T) {
	if 1 == math.MaxInt {
		t.Log("no")
	} else if 1 < math.MaxInt {
		t.Log("second if branch!")
	} else {
		t.Log("final kamen ride!")
	}
}
