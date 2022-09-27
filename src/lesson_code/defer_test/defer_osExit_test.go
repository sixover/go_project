package defer_test

import (
	"testing"
)

func TestDeferOsExit(t *testing.T) {
	defer func() {
		t.Log("i am defer print!!!!")
		return
	}()
	panic("its a panic")
}
