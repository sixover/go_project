package stack_and_queue

type Queue struct {
	size  int
	limit int
	pushi int
	popi  int
	arr   []int
}

func getQueue(i int) *Queue {
	return &Queue{
		size:  0,
		limit: i,
		pushi: 0,
		popi:  0,
		arr:   make([]int, i),
	}
}

func pop(q *Queue) int {
	if q.size == 0 {
		return -1
	}
	ret := q.arr[q.popi]
	q.popi = (q.popi + 1) % q.limit
	q.size--
	return ret
}
func push(q *Queue, i int) {
	if q.size == q.limit {
		return
	}
	q.arr[q.pushi] = i
	q.pushi = (q.pushi + 1) % q.limit
	q.size++
}
