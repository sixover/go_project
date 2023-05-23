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
type class []student

func (c *class) Less(i, j int) bool {
	return (*c)[i].age < (*c)[j].age
}
func (c *class) Len() int {
	return len(*c)
}
func (c *class) Swap(i, j int) {
	(*c)[i], (*c)[j] = (*c)[j], (*c)[i]
}

type aa interface {
	what()
}

func (c *class) what() {
	fmt.Println((*c))
}

func TestSortBySelf(t *testing.T) {
	rand.Seed(time.Now().Unix())
	a := class{}
	for i := 0; i < 10; i++ {
		a = append(a, student{name: strconv.Itoa(rand.Intn(100)), age: rand.Intn(100)})
	}
	fmt.Println(a)
	sort.Sort(&a)
	fmt.Println(a)
	pos := sort.Search(a.Len(), func(i int) bool {
		return a[i].age >= 60
	})
	fmt.Println(pos)
	test_interface := aa(&a)
	test_interface.what()
	//b:=list.New()
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
func TestMap(t *testing.T) {
	mmap := make(map[int]int)
	v, ok := mmap[1]
	fmt.Println(v, ok)
	_, _ = mmap[2]
}
