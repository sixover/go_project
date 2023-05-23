package quick_sort

import (
	"math/rand"
	"time"
)

/*
快速排序算法
其实就是以数组中最后一个数，作为X的荷兰国旗问题的算法
用函数进行递归，每一次递归将其分为<X和=X，>X的区域，然后返回=X的下标。
例如arr=[1,4,2,3,3]，第一次排序F(0,len(arr)-1)排成了[1,2,3,3,4]，然后返回了[3,3]的下标[2,3]
下次递归排序就是F(0,1)，F(4,len(arr)-1)，以此类推
*/
func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	rand.Seed(time.Now().Unix())
	progress(arr, 0, len(arr)-1)
	return
}
func progress(arr []int, l, r int) {
	if l >= r {
		return
	}
	swap(arr, l+rand.Intn(r-l+1), r)
	left, right := sort(arr, l, r)
	progress(arr, l, left-1)
	progress(arr, right+1, r)
}
func sort(arr []int, l, r int) (int, int) {
	if l == r {
		return l, r
	}
	index := l
	left := l - 1
	// 以数组最后一位为X，把arr[l,r-1]位置看作数组进行排序，最后将>X最左边的数与其swap
	right := r
	for right > index {
		if arr[index] == arr[r] {
			index++
		} else if arr[index] < arr[r] {
			swap(arr, index, left+1)
			index++
			left++
		} else {
			swap(arr, index, right-1)
			right--
		}
	}
	swap(arr, right, r)
	return left + 1, right
}
func qui(arr []int) {
	if len(arr) <= 1 {
		return
	}
	pro(arr, 0, len(arr)-1)
}
func pro(arr []int, i, j int) {
	if i >= j {
		return
	}
	l, r := sssort(arr, i, j, arr[j])
	pro(arr, i, l-1)
	pro(arr, r+1, j)
}
func sssort(arr []int, l, r, x int) (int, int) {
	left := l - 1
	right := r + 1
	index := l
	for index < right {
		if arr[index] < x {
			arr[left+1], arr[index] = arr[index], arr[left+1]
			left++
			index++
		} else if arr[index] == x {
			index++
		} else {
			arr[right-1], arr[index] = arr[index], arr[right-1]
			right--
		}
	}
	return left + 1, right - 1
}

func quickStart(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	ret := prog(arr, 0, len(arr)-1, false)
	return ret
}

func prog(arr []int, i, j int, leftover bool) []int {
	if i >= j {
		if leftover {
			return []int{arr[i]}
		} else {
			return []int{arr[j]}
		}

	}
	temp, l, r := ssort(arr, i, j, arr[j])
	left := prog(temp, i, l-1, true)
	right := prog(temp, r+1, j, false)
	ret := []int{}
	ret = append(ret, left...)
	ret = append(ret, right...)
	return ret
}

func ssort(arr []int, l, r int, x int) ([]int, int, int) {
	if l == r {
		return []int{arr[l]}, l, r
	}
	left := l - 1
	right := r + 1
	index := l
	for index < right {
		if arr[index] < x {
			arr[left+1], arr[index] = arr[index], arr[left+1]
			index++
			left++
		} else if arr[index] == x {
			index++
		} else {
			arr[right-1], arr[index] = arr[index], arr[right-1]
			right--
		}
	}
	return arr, left + 1, right - 1
}
