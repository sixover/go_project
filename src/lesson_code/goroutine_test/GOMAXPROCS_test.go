package goroutine

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGOMAXPROCS(t *testing.T) {
	runtime.GOMAXPROCS(2)
	for i := 0; i < 50; i++ {
		go fmt.Print(0)
		fmt.Print(1)
	}
}
