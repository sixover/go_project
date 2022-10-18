package heap_sort

/*
	将一个切片调整为一个大根堆或者小根堆，不需要pop和push函数，只需要这两个就足够了
	这里写大根堆的方法
*/
func heapInsert(arr []int, index int) {
	for index >= 0 && arr[index] > arr[(index-1)/2] {
		swap(arr, index, (index-1)/2)
		index = (index - 1) / 2
	}
}
func heapAdjust(arr []int, index, size int) {
	left := index*2 + 1
	for left < size {
		large := left
		if left+1 < size {
			if arr[left+1] > arr[left] {
				large = left + 1
			}
		}
		if arr[index] >= arr[large] {
			break
		}
		swap(arr, index, large)
		index = large
		left = index*2 + 1
	}
}
