package xor_test

import (
	"fmt"
	"testing"
)

func TestOnlyKtimes(t *testing.T) {
	arr := []int{1, 1, 1, 1, 1, 23, 23, 23, 23, 23, 4, 4, 4, 12, 12, 12, 12, 12, 5, 5, 5, 5, 5, 21, 21, 21, 21, 21}
	res := findNum(arr, 3, 5)
	fmt.Println(res)
}

func findNum(arr []int, k, m int) int {
	bitNum := make([]int, 64)
	for _, res := range arr {
		for i := 0; i < 64; i++ {
			if (res>>i)&1 == 1 {
				bitNum[i]++
			}
		}
	}
	ret := 0
	for i := 0; i < 64; i++ {
		//if bitNum[i]%m != 0 {
		//	//ret=ret+int(math.Pow(2, float64(i)))
		//	//ret=ret+(1<<i)
		//	ret |= 1 << i
		//}
		if bitNum[i]%m == k {
			//ret=ret+int(math.Pow(2, float64(i)))
			//ret=ret+(1<<i)
			ret |= 1 << i
		}
	}
	return ret
}
