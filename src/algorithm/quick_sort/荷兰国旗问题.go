package quick_sort

func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

/*
荷兰国旗问题，即给一个无序的数组分组，给定一个数X，将数组分为，左边是小于等于X的数，右边是大于X的数
要求时间复杂度为O（N），空间复杂度为O（1）
例如，arr=[1,234,1,235,3,452,34],X=7
分组后为
arr=[1,1,3,234,235,452,34]
*/

/*
解题思路，从左划分一个小于区，每次找到一个小于等于X的数，将其划分到小于区内，然后小于区扩张
*/
func flagSimpleSort(arr []int, x int) []int {
	if len(arr) == 0 {
		return nil
	}
	if len(arr) == 1 {
		return arr
	}
	index := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] <= x {
			swap(arr, i, index+1)
			index++
		}
	}
	return arr
}

func itsTest(arr []int, x int) []int {
	index := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] <= x {
			arr[index+1], arr[i] = arr[i], arr[index+1]
			index++
		}
	}
	return arr
}

/*
荷兰国旗问题2：
即给一个无序的数组分组，给定一个数X，将数组分为，左边是小于X的数，中间是等于X的数，右边是大于X的数
要求时间复杂度为O（N），空间复杂度为O（1）
例如，arr=[1,234,7,1,235,7,3,452,34],X=7
分组后为
arr=[1,1,3,7,7,234,235,452,34]
*/

/*
解题思路，左边划分一个小于区，右边划分一个大于区，从左往右遍历数组
找到一个小于X的数，划分到小于区中，找到一个大于X的数，划分到大于区中，找到一个等于X的数，跳过
如果不输入X，以数组最后一个元素作为X来划分，那么可以先以上面的思路划分[0-(len(arr)-1)]的数，然后将最后一个数与大于区最左边的数做交换
*/

func flagComplexSort(arr []int, x int) []int {
	if len(arr) == 0 {
		return nil
	}
	if len(arr) == 1 {
		return arr
	}
	leftIndex := -1
	rightIndex := len(arr)
	//for i:=0;i<len(arr);i++{
	//	for rightIndex>i&&arr[i]>x{
	//		swap(arr,i,rightIndex-1)
	//		rightIndex--
	//	}
	//	if rightIndex==i{
	//		break
	//	}
	//	if arr[i]<x{
	//		swap(arr,i,leftIndex+1)
	//		leftIndex++
	//	}
	//}
	index := 0
	for rightIndex > index {
		if arr[index] == x {
			index++
		} else if arr[index] < x {
			swap(arr, index, leftIndex+1)
			index++
			leftIndex++
		} else {
			swap(arr, index, rightIndex-1)
			rightIndex--
		}
	}
	return arr
}

func flag2QS(arr []int, x int) []int {
	left := -1
	right := len(arr)
	index := 0
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
	return arr
}
