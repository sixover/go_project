package merge_sort

func countRangeSum(nums []int, lower int, upper int) int {
	presum := make([]int, len(nums))
	presum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		presum[i] = presum[i-1] + nums[i]
	}
	ret := progress(presum, 0, len(presum)-1, lower, upper)
	return ret
}

func progress(nums []int, l, r, lower, upper int) int {
	if l == r {
		if nums[l] >= lower && nums[l] <= upper {
			return 1
		} else {
			return 0
		}
	}
	mid := l + (r-l)/2
	left := progress(nums, l, mid, lower, upper)
	right := progress(nums, mid+1, r, lower, upper)
	ret := left + right + merge(nums, l, mid, r, lower, upper)
	return ret
}
func merge(nums []int, l, m, r, lower, upper int) int {
	count := 0
	leftIndexMin := l
	leftIndexMax := l
	rightIndex := m + 1
	for rightIndex <= r {
		for leftIndexMin <= m && (nums[rightIndex]-nums[leftIndexMin]) > upper {
			leftIndexMin++
		}
		for leftIndexMax <= m && (nums[rightIndex]-nums[leftIndexMax]) >= lower {
			leftIndexMax++
		}
		count += (leftIndexMax - leftIndexMin)
		rightIndex++
	}

	mid := m + 1
	left := l
	index := 0
	temp := make([]int, r-l+1)
	for left <= m && mid <= r {
		if nums[left] < nums[mid] {
			temp[index] = nums[left]
			left++
			index++
		} else {
			temp[index] = nums[mid]
			mid++
			index++
		}
	}
	for left <= m {
		temp[index] = nums[left]
		left++
		index++
	}
	for mid <= r {
		temp[index] = nums[mid]
		mid++
		index++
	}
	for i := l; i <= r; i++ {
		nums[i] = temp[i-l]
	}
	return count
}
