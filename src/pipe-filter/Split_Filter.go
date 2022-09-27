package pipe_filter

import (
	"errors"
	"strings"
)

type Split struct {
	fenge string
}

func GetSplitObject(fenge string) *Split {
	return &Split{fenge: fenge}
}

func (s *Split) Progress(inputData Input) (Output, error) {
	str, ok := inputData.(string)
	if !ok {
		return nil, errors.New("the input is not string, please input the correct struct!")
	}
	parts := strings.Split(str, s.fenge)
	return parts, nil
}
