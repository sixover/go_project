package merge_sort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	a := []int{7, 8, 1, 888, 54, 654, 3, 213, 1, 65, 487}
	b := mergeProcess1(a, 0, len(a)-1)
	fmt.Println(b)
	c := []int{7, 8, 1, 888, 54, 654, 3, 213, 1, 65, 487}
	mergeProcess2(c, 0, len(c)-1)
	fmt.Println(c)
	d := []int{7, 8, 1, 888, 54, 654, 3, 213, 1, 65, 487}
	mergeNonProcess(d)
	fmt.Println(d)
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
		mergeSize *= 2
	}
}
