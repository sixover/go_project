package stack_and_queue

type queueByStack struct {
	a stack
	b stack
}

type stack struct {
	arr   []int
	index int
}

func (s *stack) push(i int) {
	s.arr = append(s.arr, i)
	s.index++
}
func (s *stack) pop() int {
	if s.index == 0 {
		return -1
	}
	ret := s.arr[s.index-1]
	s.arr = s.arr[:s.index-1]
	s.index--
	return ret
}

func getQueueByStack() *queueByStack {
	return &queueByStack{
		a: stack{},
		b: stack{},
	}
}

func (q *queueByStack) push(i int) {
	q.a.push(i)
}
func (q *queueByStack) pop() int {
	if q.b.index == 0 && q.a.index == 0 {
		return -1
	}
	if q.b.index == 0 {
		for {
			i := q.a.pop()
			if i == -1 {
				break
			}
			q.b.push(i)
		}
	}
	return q.b.pop()
}
