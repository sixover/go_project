package count_sort

//arr中元素范围[0-200]
func countSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	maxNum := arr[0]
	for _, num := range arr {
		if num > maxNum {
			maxNum = num
		}
	}
	//0-maxnum
	tempArr := make([]int, maxNum+1)
	for _, num := range arr {
		tempArr[num]++
	}

	i := 0
	for j := 0; j < len(tempArr); j++ {
		for tempArr[j] > 0 {
			arr[i] = j
			i++
			tempArr[j]--
		}
	}
	return
}
