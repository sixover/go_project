package pipe_filter

import (
	"errors"
	"fmt"
)

type Run struct {
	str    string
	Filter *[]filter
}

func GetRunObject(str string, Filters ...filter) *Run {
	return &Run{
		str:    str,
		Filter: &Filters,
	}
}

//一个是...的数据形式，一个是*r.Filter
func (r *Run) Progress(input Input) (Output, error) {
	fmt.Printf("the r.Filter type is : %T\n", r.Filter)
	fmt.Printf("the *r.Filter type is : %T\n", *r.Filter)
	for _, unit := range *r.Filter {
		ret, err := unit.Progress(input)
		if err != nil {
			return nil, errors.New("one unit filter run wrong!")
		}
		input = ret
	}
	return input, nil
}
