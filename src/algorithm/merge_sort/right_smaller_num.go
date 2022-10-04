package merge_sort

/*
	其实这个过程可以相对来看，一个是找右边比较小的数的个数，这个需要在两个组之间从后往前merge
	也可以找左边比较大的数的个数，依旧是可以在两个组之间从前往后merge
	这里两个都写一下
*/
func findNumOfRightSmaller1(arr []int, l, r int) int {
	if l == r {
		return 0
	}
	mid := l + (r-l)/2
	left := findNumOfRightSmaller1(arr, l, mid)
	right := findNumOfRightSmaller1(arr, mid+1, r)
	ret := left + right + mergeRightSmaller1(arr, l, mid, r)
	return ret
}

//从后往前merge
func mergeRightSmaller1(arr []int, l, m, r int) int {
	tempSlice := make([]int, r-l+1)
	index := r - l
	Mid := m
	R := r
	count := 0
	for Mid >= l && R >= m+1 {
		if arr[Mid] > arr[R] {
			tempSlice[index] = arr[Mid]
			count += R - m
			index--
			Mid--
		} else {
			tempSlice[index] = arr[R]
			index--
			R--
		}
	}
	for Mid >= l {
		tempSlice[index] = arr[Mid]
		index--
		Mid--
	}
	for R >= m+1 {
		tempSlice[index] = arr[R]
		index--
		R--
	}
	for i := l; i <= r; i++ {
		arr[i] = tempSlice[i-l]
	}
	return count
}

func findNumOfRightSmaller2(arr []int, l, r int) int {
	if l == r {
		return 0
	}
	mid := l + (r-l)/2
	left := findNumOfRightSmaller2(arr, l, mid)
	right := findNumOfRightSmaller2(arr, mid+1, r)
	ret := left + right + mergeRightSmaller2(arr, l, mid, r)
	return ret
}
func mergeRightSmaller2(arr []int, l, m, r int) int {
	tempSlice := make([]int, r-l+1)
	index := 0
	L := l
	Mid := m + 1
	count := 0
	for L <= m && Mid <= r {
		if arr[L] > arr[Mid] {
			tempSlice[index] = arr[Mid]
			count += m - L + 1
			Mid++
			index++
		} else {
			tempSlice[index] = arr[L]
			L++
			index++
		}
	}
	for L <= m {
		tempSlice[index] = arr[L]
		index++
		L++
	}
	for Mid <= r {
		tempSlice[index] = arr[Mid]
		index++
		Mid++
	}
	for i := l; i <= r; i++ {
		arr[i] = tempSlice[i-l]
	}
	return count
}
