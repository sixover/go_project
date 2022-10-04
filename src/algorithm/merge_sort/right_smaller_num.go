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
func mergeRightSmaller1(arr []int, l, m, r int) int {

}
