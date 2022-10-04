package merge_sort

func findNumOfRightSmallerDoubleNum(arr []int, l, r int) int {
	if l == r {
		return 0
	}
	mid := l + (r-l)/2
	left := findNumOfRightSmallerDoubleNum(arr, l, mid)
	right := findNumOfRightSmallerDoubleNum(arr, mid+1, r)
	ret := left + right + mergeRightSmallerDouble(arr, l, mid, r)
	return ret
}
func mergeRightSmallerDouble(arr []int, l, m, r int) int {
	//O(N)计算该merge中符合条件的数有多少
	//countL := l
	//countR := m + 1
	//count:=0
	//for countL <= m && countR <= r {
	//	if arr[countL] <= (arr[countR] * 2) {
	//		countL++
	//		count+=countR-m-1
	//	}else{
	//		countR++
	//	}
	//}
	//if countR>r{
	//	count+=(countR-m-1)*(m-countL+1)
	//}
	count := 0
	countR := m + 1
	for i := l; i <= m; i++ {
		for countR <= r && arr[i] > (arr[countR]*2) {
			countR++
		}
		count += countR - m - 1
	}

	//正常merge
	tempSlice := make([]int, r-l+1)
	index := 0
	L := l
	Mid := m + 1
	for L <= m && Mid <= r {
		if arr[L] > arr[Mid] {
			tempSlice[index] = arr[Mid]
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
