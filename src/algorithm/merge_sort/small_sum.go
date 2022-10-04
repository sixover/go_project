package merge_sort

func small_sum(arr []int, l, r int) int {
	if len(arr) < 2 {
		return 0
	}
	if l == r {
		return 0
	}
	mid := l + (r-l)/2
	left := small_sum(arr, l, mid)
	right := small_sum(arr, mid+1, r)
	ret := smallSumMerge(arr, l, mid, r) + left + right
	return ret
}
func smallSumMerge(arr []int, l, m, r int) int {
	tempSlice := make([]int, r-l+1)
	index := 0
	L := l
	Mid := m + 1
	smallCount := 0
	for L <= m && Mid <= r {
		if arr[L] < arr[Mid] {
			tempSlice[index] = arr[L]
			smallCount += (r - Mid + 1) * arr[L]
			L++
			index++
		} else {
			tempSlice[index] = arr[Mid]
			Mid++
			index++
		}
	}
	for L <= m {
		tempSlice[index] = arr[L]
		L++
		index++
	}
	for Mid <= r {
		tempSlice[index] = arr[Mid]
		Mid++
		index++
	}
	for i := l; i <= r; i++ {
		arr[i] = tempSlice[i-l]
	}
	return smallCount
}
