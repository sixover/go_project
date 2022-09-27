package pipe_filter

import "errors"

type Sum struct{}

func GetSumObject() *Sum {
	return &Sum{}
}
func (S *Sum) Progress(input Input) (Output, error) {
	parts, ok := input.([]int)
	if !ok {
		return nil, errors.New("the input is not a []int, wrong input!")
	}
	sum := 0
	for _, res := range parts {
		sum += res
	}
	return sum, nil
}
