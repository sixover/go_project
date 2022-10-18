package heap_sort

type heapSmall struct {
	heap []int
	size int
}

func (h *heapSmall) heapInsert(index int) {
	for index >= 0 && h.heap[(index-1)/2] > h.heap[index] {
		swap(h.heap, (index-1)/2, index)
		index = (index - 1) / 2
	}
}

func (h *heapSmall) heapAdjust(index int) {
	left := index*2 + 1
	for left < h.size {
		small := left
		if left+1 < h.size {
			if h.heap[left] < h.heap[left+1] {
				small = left
			} else {
				small = left + 1
			}
		}
		if h.heap[index] <= h.heap[small] {
			break
		}
		swap(h.heap, index, small)
		index = small
		left = index*2 + 1
	}
}

func (h *heapSmall) push(i int) {
	h.heap = append(h.heap, i)
	h.size++
	h.heapInsert(h.size - 1)
}
func (h *heapSmall) pop() int {
	ret := h.heap[0]
	swap(h.heap, 0, h.size-1)
	h.size--
	h.heap = h.heap[:h.size]
	h.heapAdjust(0)
	return ret
}
func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}
