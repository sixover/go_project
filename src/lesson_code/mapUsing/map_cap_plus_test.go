package mapUsing

import (
	"fmt"
	"testing"
)

func TestMapCapPlus(t *testing.T) {
	tMap := make(map[int]int, 4)
	for i := 0; i < 10; i++ {
		tMap[i] = i
		fmt.Println(len(tMap))
	}
}

func TestNilValueAndTrueValue(t *testing.T) {
	tMap := map[int]int{}
	fmt.Println(tMap[1])
	tMap[2] = 0
	fmt.Println(tMap[2])
	if value, ok := tMap[1]; ok {
		fmt.Println("the tMap[1] value is:", value)
	} else {
		fmt.Println("the tMap[1] value not exist!")
	}
}
