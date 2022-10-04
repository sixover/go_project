package merge_sort

func findNumOfRightBigger(arr []int, l, r int) int {
	if l == r {
		return 0
	}
	mid := l + (r-l)/2
	left := findNumOfRightBigger(arr, l, mid)
	right := findNumOfRightBigger(arr, mid+1, r)
	ret := mergeRightBigger(arr, l, mid, r) + left + right
	return ret
}
func mergeRightBigger(arr []int, l, m, r int) int {
	tempSlice := make([]int, r-l+1)
	index := 0
	Mid := m + 1
	L := l
	count := 0
	for L <= m && Mid <= r {
		if arr[L] < arr[Mid] {
			count += r - Mid + 1
			tempSlice[index] = arr[L]
			L++
		} else {
			tempSlice[index] = arr[Mid]
			Mid++
		}
		index++
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
	return count
}
