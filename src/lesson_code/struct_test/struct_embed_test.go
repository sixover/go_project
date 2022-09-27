package struct_test

import "testing"

type One struct {
	x int
	y int
}

func (o *One) sqrt() int {
	return o.x*o.x + o.y*o.y
}

type TwoO struct {
	o One
	z int
}

func (t *TwoO) sqrt() int {
	return t.o.x * t.o.y * t.z
}

type TwoT struct {
	One
	z int
}

func (t *TwoT) sqrt() int {
	return t.x * t.y * t.z
}

type TwoS struct {
	o  One
	oo One
	z  int
}

func (t *TwoS) sqrt() int {
	return t.o.sqrt() + t.oo.sqrt() + t.z
}

type TwoF struct {
	o One
	One
	z int
}

func TestEmbed(t *testing.T) {
	to := &TwoO{
		o: One{
			x: 1,
			y: 2,
		},
		z: 3,
	}
	tt := &TwoT{
		One: One{
			x: 1,
			y: 2,
		},
		z: 3,
	}
	ts := &TwoS{
		o: One{
			x: 1,
			y: 2,
		},
		oo: One{
			x: 1,
			y: 2,
		},
		z: 3,
	}
	t.Log("to's One sqrt is : ", to.o.sqrt())
	t.Log("to's sqrt is : ", to.sqrt())
	t.Log("tt's one sqrt is : ", tt.One.sqrt())
	t.Log("tt's sqrt is : ", tt.sqrt())
	t.Log("ts's double One sqrt is : ", ts.sqrt())
}
