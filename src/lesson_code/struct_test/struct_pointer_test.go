package struct_test

import (
	"fmt"
	"testing"
	"unsafe"
)

type PointTest struct {
	X int
	Y int
}

func (p PointTest) PrintPointAddressOne() string {
	fmt.Printf("the Aaddress of struct is : %x\n", unsafe.Pointer(&p))
	return fmt.Sprintf("the Baddress of struct is : %x\n", unsafe.Pointer(&p))
}

func (p *PointTest) PrintPointAddressTwo() string {
	fmt.Printf("the Aaddress of struct is : %x\n", unsafe.Pointer(p))
	return fmt.Sprintf("the Baddress of struct is : %x\n", unsafe.Pointer(p))
}

func TestStructAddress(t *testing.T) {
	p := &PointTest{
		X: 1,
		Y: 2,
	}
	t.Log(unsafe.Pointer(p))
	t.Log(p.PrintPointAddressOne())
	t.Log(p.PrintPointAddressTwo())
}
