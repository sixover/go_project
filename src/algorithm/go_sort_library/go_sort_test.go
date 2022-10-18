package go_sort_library

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"testing"
	"time"
)

type student struct {
	name string
	age  int
}

func TestSortLibrary(t *testing.T) {
	a := []int{1, 3, 53, 2, 23, 46, 2, 346, 347, 22}
	fmt.Println(sort.IntsAreSorted(a))
	sort.Ints(a)
	fmt.Println(a)
	fmt.Println(sort.IntsAreSorted(a))
	res := sort.SearchInts(a, 7)
	fmt.Println(res)

	b := []int{1, 3, 53, 2, 23, 46, 2, 346, 347, 22}
	c := sort.IntSlice(b)
	c.Sort()
	fmt.Println(c)
	fmt.Println(c.Search(7))
	sort.Sort(sort.Reverse(c))
	fmt.Println(c)
}
func TestSliceSort(t *testing.T) {
	rand.Seed(time.Now().Unix())
	a := []student{}
	for i := 0; i < 10; i++ {
		a = append(a, student{name: strconv.Itoa(rand.Intn(100)), age: rand.Intn(100)})
	}
	fmt.Println(a)
	sort.Slice(a, func(i, j int) bool {
		return a[i].age < a[j].age
	})
	fmt.Println(a)
	fmt.Println(sort.SliceIsSorted(a, func(i, j int) bool {
		return a[i].age < a[j].age
	}))
}
func TestSearch(t *testing.T) {
	a := []int{1, 3, 53, 2, 23, 46, 2, 346, 347, 22}
	sort.Ints(a)
	fmt.Println(a)
	res := sort.Search(len(a), func(i int) bool {
		return a[i] > 3
	})
	fmt.Println(res)
	res = sort.Search(len(a), func(i int) bool {
		return a[i] < 23
	})
	fmt.Println(res)
}
