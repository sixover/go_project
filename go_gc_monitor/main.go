package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"time"
)

func printGCStates() {
	t := time.NewTicker(time.Second)
	s := debug.GCStats{}
	for {
		select {
		case <-t.C:
			debug.ReadGCStats(&s)
			fmt.Printf("gc %d last@%v, PauseTotal %v\n", s.NumGC, s.LastGC, s.PauseTotal)
		}
	}
}
func printMemStates() {
	t := time.NewTicker(time.Second)
	s := runtime.MemStats{}
	for {
		select {
		case <-t.C:
			runtime.ReadMemStats(&s)
			fmt.Printf("gc %d last@%v, next_heap_size@%vMB\n", s.NumGC, time.Unix(int64(time.Duration(s.LastGC).Seconds()), 0), s.NextGC/(1<<20))
			//fmt.Printf("gc %d last@%v, next_heap_size@%vMB\n", s.NumGC, s.LastGC, s.NextGC/(1<<20))
		}
	}
}

func main() {
	//go printGCStates()
	go printMemStates()
	for i := 0; i < 100000; i++ {
		_ = make([]byte, 1<<28)
	}

}
