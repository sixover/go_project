package string_test

import "testing"

func TestString(t *testing.T) {
	s := "中"
	t.Log(len(s))
	ss := []rune(s)
	t.Log(len(ss))
	t.Logf("中 unicode %x", ss[0])
	t.Logf("中 utf-8 %x", s)
	s = "中文"
	t.Log(len(s))
	ss = []rune(s)
	t.Log(len(ss))
}
