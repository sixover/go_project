package recursion

func getIntNumMax(arr []int, l, r int) int {
	if l == r {
		return arr[l]
	}
	mid := l + (r-l)/2
	left := getIntNumMax(arr, l, mid)
	right := getIntNumMax(arr, mid+1, r)
	if left > right {
		return left
	} else {
		return right
	}
}
