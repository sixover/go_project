package merge_sort

func mergedigui(arr []int, l, r int) []int {
	if l == r {
		return []int{arr[l]}
	}
	mid := l + (r-l)/2
	left := mergedigui(arr, l, mid)
	right := mergedigui(arr, mid+1, r)
	ret := mergeing(left, right)
	return ret
}
func mergeing(left []int, right []int) []int {
	index := 0
	l := 0
	r := 0
	ret := make([]int, len(left)+len(right))
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			ret[index] = left[l]
			l++
		} else {
			ret[index] = right[r]
			r++
		}
		index++
	}
	for l < len(left) {
		ret[index] = left[l]
		l++
		index++
	}
	for r < len(right) {
		ret[index] = right[r]
		r++
		index++
	}
	return ret
}

func mergeProcess1(arr []int, l, r int) []int {
	if l == r {
		return []int{arr[l]}
	}
	mid := l + (r-l)/2
	left := mergeProcess1(arr, l, mid)
	right := mergeProcess1(arr, mid+1, r)
	ret := merge1(left, right)
	return ret
}
func merge1(left, right []int) []int {
	ret := make([]int, len(left)+len(right))
	index := 0
	l := 0
	r := 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			ret[index] = left[l]
			l++
		} else {
			ret[index] = right[r]
			r++
		}
		index++
	}
	for l != len(left) {
		ret[index] = left[l]
		l++
		index++
	}
	for r != len(right) {
		ret[index] = right[r]
		r++
		index++
	}
	return ret
}
func mergeProcess2(arr []int, l, r int) {
	if l == r {
		return
	}
	mid := l + (r-l)/2
	mergeProcess2(arr, l, mid)
	mergeProcess2(arr, mid+1, r)
	merge2(arr, l, mid, r)
}
func merge2(arr []int, l, m, r int) {
	temp := make([]int, r-l+1)
	index := 0
	templ := l
	tempMid := m + 1
	for templ <= m && tempMid <= r {
		if arr[templ] < arr[tempMid] {
			temp[index] = arr[templ]
			templ++
		} else {
			temp[index] = arr[tempMid]
			tempMid++
		}
		index++
	}
	for templ <= m {
		temp[index] = arr[templ]
		templ++
		index++
	}
	for tempMid <= r {
		temp[index] = arr[tempMid]
		tempMid++
		index++
	}
	for i := l; i <= r; i++ {
		arr[i] = temp[i-l]
	}
}
func mergeNonProcess(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	N := len(arr)
	//步长
	mergeSize := 1
	//如果步长大于N，则排序完成
	for mergeSize < N {
		L := 0
		//如果第一轮L>N，那么一轮排序完成，需要调整步长
		for L < N {
			//首先确定每轮的左边一组，即L-M
			Mid := L + mergeSize - 1
			//如果Mid>=N，则左边一组缺省，或者无右边一组，跳出循环，这一轮排序结束
			if Mid >= N {
				break
			}
			//若是不缺省，则左边一组与右边一组进行merge，注意右边一组的边界
			R := 0
			if Mid+mergeSize < N-1 {
				R = Mid + mergeSize
			} else {
				R = N - 1
			}
			merge2(arr, L, Mid, R)
			//进行下一组的排序
			L = R + 1
		}
		//这一步非常重要，当N非常接近int的最大值时，如果此时不跳出，等待执行mergeSize *= 2
		//那么执行完毕后mergeSize可能会造成溢出，从而变成负数，则循环出错
		if mergeSize > N/2 {
			break
		}
		mergeSize *= 2
	}
}
