package error_test

import (
	"errors"
	"fmt"
	"testing"
)

func Fibonaci(n int) ([]int, error) {
	temp := []int{1, 1}
	if n < 2 {
		return nil, errors.New("cannot less than 2")
	}
	if n > 100 {
		return nil, errors.New("cannot more than 100")
	}
	for i := 2; i < n; i++ {
		temp = append(temp, temp[i-2]+temp[i-1])
	}
	return temp, nil
}

func TestErrorUsing(t *testing.T) {
	list, error := Fibonaci(19)
	if error != nil {
		t.Log(error)
		return
	}
	fmt.Println(list)
}
