package nums_and_slice

import (
	"fmt"
	"testing"
)

func TestSliceInit(t *testing.T) {
	var s []int
	s = append(s, 1)
	ss := make([]int, 3, 5)
	ss = append(ss, 23)
	fmt.Println(s)
	fmt.Println(ss)
}
