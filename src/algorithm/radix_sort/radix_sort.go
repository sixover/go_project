package radix_sort

import (
	"container/list"
	"math"
)

//函数名不知道怎么说了，功能是获得数组中最大值是几位数
func maxDigitBitInTen(arr []int, L, R int) int {
	maxNum := arr[L]
	for i := L + 1; i <= R; i++ {
		if arr[i] > maxNum {
			maxNum = arr[i]
		}
	}

	count := 0
	for maxNum != 0 {
		maxNum = maxNum / 10
		count++
	}
	return count
}

//获取在个十百千位上的数
func getDigitNum(num, wei int) int {
	return num / int(math.Pow(10, float64(wei-1))) % 10
}

//在arr数组中的L，R区间应用基数排序
func radixSort(arr []int, L, R int) {
	digitNum := maxDigitBitInTen(arr, L, R)
	tempRes := make([]int, R-L+1)
	for i := 1; i <= digitNum; i++ {
		count := make([]int, 10)

		for j := L; j <= R; j++ {
			res := getDigitNum(arr[j], i)
			count[res]++
		}
		//转换成前缀和数组
		for j := 1; j < len(count); j++ {
			count[j] = count[j-1] + count[j]
		}

		//从右往左遍历原数组
		for j := R; j >= L; j-- {
			res := getDigitNum(arr[j], i)
			tempRes[count[res]-1] = arr[j]
			count[res]--
		}

		for j := L; j <= R; j++ {
			arr[j] = tempRes[j-L]
		}
	}
}

func radixSort2(arr []int, L, R int) {
	dig := maxDigitBitInTen(arr, L, R)
	bucketForSort := make([]*list.List, 10)
	for i := 0; i < len(bucketForSort); i++ {
		bucketForSort[i] = list.New()
	}
	for i := 1; i <= dig; i++ {
		for j := L; j <= R; j++ {
			res := getDigitNum(arr[j], i)
			bucketForSort[res].PushFront(arr[j])
		}

		index := L
		for j := 0; j < len(bucketForSort); j++ {
			for bucketForSort[j].Len() != 0 {
				temp := bucketForSort[j].Back()
				bucketForSort[j].Remove(temp)
				arr[index] = temp.Value.(int)
				index++
			}
		}
	}
}
