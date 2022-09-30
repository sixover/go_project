package stack_and_queue

type Stack struct {
	index int
	arr   []int
}

func getStackInit() *Stack {
	return &Stack{
		index: 0,
		arr:   []int{},
	}
}
func (s *Stack) push(i int) {
	s.arr = append(s.arr, i)
	s.index++
}
func (s *Stack) pop() int {
	if s.index == 0 {
		return -1
	}
	ret := s.arr[s.index-1]
	s.index--
	if s.index == 0 {
		s.arr = []int{}
	} else {
		s.arr = s.arr[:s.index]
	}
	return ret
}
