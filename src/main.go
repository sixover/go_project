package main

import (
	"fmt"
	pf "go_project/src/pipe-filter"
)

func main() {
	inputStr := "1,2,3"
	splitFilter := pf.GetSplitObject(",")
	toInt := pf.GetToIntObject()
	sum := pf.GetSumObject()
	allFilter := pf.GetRunObject("filterOne", splitFilter, toInt, sum)
	ret, err := allFilter.Progress(inputStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret)
}
