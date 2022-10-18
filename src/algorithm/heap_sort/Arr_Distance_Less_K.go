package heap_sort

/*
	这种情况就是说，已知一个几乎有序的数组。
	几乎有序是指，如果把数组排好顺序的话，每个元素移动的距离一定不超过k，并且k相对于数组长度来说是比较小的。
	请选择一个合适的排序策略，对这个数组进行排序。
*/
/*
	这道题就是，对于数组中0这个位置来说，只有前k+1个数是可能在这个位置上的
	所以首先对前k+1建堆
	然后将堆顶拿出来放在数组0的位置
	然后将第k+2个数加入到这个堆中，将堆顶的数取出来放到数组1的位置
	以此类推
	时间复杂度O（N*LogK）
*/
func sortedArrDistanceLessK(arr []int, k int) {
	heap := new(heapSmall)
	index := 0
	for index <= k {
		heap.push(arr[index])
		index++
	}
	i := 0
	for index < len(arr) {
		arr[i] = heap.pop()
		heap.push(arr[index])
		index++
		i++
	}
	for len(heap.heap) > 0 {
		arr[i] = heap.pop()
		i++
	}
}
