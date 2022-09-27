package pipe_filter

import (
	"errors"
	"strconv"
)

type ToInt struct{}

func GetToIntObject() *ToInt {
	return &ToInt{}
}

func (t *ToInt) Progress(inputData Input) (Output, error) {
	ints, ok := inputData.([]string)
	if !ok {
		return nil, errors.New("this is not a []string, input a correct struct!")
	}
	ret := make([]int, 1)
	for _, res := range ints {
		i, err := strconv.Atoi(res)
		if err != nil {
			return nil, errors.New("strconv.Atoi has wrong!")
		}
		ret = append(ret, i)
	}
	return ret, nil
}
