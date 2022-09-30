package stack_and_queue

type minStack struct {
	data  []int
	min   []int
	index int
}

func getMinStackInit() *minStack {
	return &minStack{
		data:  nil,
		min:   nil,
		index: 0,
	}
}
func (m *minStack) push(i int) {
	m.data = append(m.data, i)
	if m.index == 0 || m.min[m.index-1] > i {
		m.min = append(m.min, i)
	} else {
		m.min = append(m.min, m.min[m.index-1])
	}
	m.index++
}
func (m *minStack) pop() int {
	if m.index == 0 {
		return -1
	}
	ret := m.data[m.index-1]
	m.min = m.min[:m.index-1]
	m.data = m.data[:m.index-1]
	m.index--
	return ret
}
func (m *minStack) getMinValue() int {
	if m.index == 0 {
		return -1
	}
	return m.min[m.index-1]
}
