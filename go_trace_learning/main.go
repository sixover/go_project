package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func mmain() {
	traceFile, err := os.Create("trace_file")
	defer func(traceFile *os.File) {
		err := traceFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(traceFile)
	if err != nil {
		return
	}
	err = trace.Start(traceFile)
	defer trace.Stop()
	if err != nil {
		return
	}
	ch := make(chan int)
	go func() {
		ch <- 42
	}()
	<-ch
}
