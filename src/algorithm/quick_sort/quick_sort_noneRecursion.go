package quick_sort

import (
	"math/rand"
	"time"
)

type qujian struct {
	L int
	R int
}

func qujianFastInit(l, r int) qujian {
	return qujian{
		L: l,
		R: r,
	}
}

type stack struct {
	stack []qujian
	index int
}

func (s *stack) push(LR qujian) {
	s.stack = append(s.stack, LR)
	s.index++
}
func (s *stack) pop() qujian {
	if s.index == 0 {
		panic("pop while index == 0")
	}
	ret := s.stack[s.index-1]
	s.index--
	s.stack = s.stack[:s.index]
	return ret
}

func quickSortNoneRecursion(arr []int) {
	if len(arr) <= 1 {
		return
	}
	s := new(stack)
	rand.Seed(time.Now().Unix())
	swap(arr, rand.Intn(len(arr)), len(arr)-1)
	left, right := sort(arr, 0, len(arr)-1)
	s.push(qujianFastInit(0, left-1))
	s.push(qujianFastInit(right+1, len(arr)-1))
	for len(s.stack) != 0 {
		q := s.pop()
		if q.L < q.R {
			swap(arr, q.L+rand.Intn(q.R-q.L+1), q.R)
			left, right = sort(arr, q.L, q.R)
			s.push(qujianFastInit(q.L, left-1))
			s.push(qujianFastInit(right+1, q.R))
		}
	}
}
