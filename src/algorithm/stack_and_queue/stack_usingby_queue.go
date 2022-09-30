package stack_and_queue

type stackByQueue struct {
	a []int
	b []int
}

func getStackByQueue() *stackByQueue {
	return &stackByQueue{
		a: nil,
		b: nil,
	}
}

func (s *stackByQueue) push(i int) {
	if len(s.a) != 0 {
		s.a = append(s.a, i)
	} else {
		s.b = append(s.b, i)
	}
}
func (s *stackByQueue) pop() int {
	if len(s.a) == 0 && len(s.b) == 0 {
		return -1
	}
	var ret int
	if len(s.a) != 0 {
		ret = s.a[len(s.a)-1]
		s.b = s.a[:len(s.a)-1]
		s.a = nil
	} else {
		ret = s.b[len(s.b)-1]
		s.a = s.b[:len(s.b)-1]
		s.b = nil
	}
	return ret
}
