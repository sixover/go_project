package readCmdParameter

import (
	"fmt"
	"os"
	"testing"
)

func TestAlgorithm(t *testing.T) {
	if len(os.Args) > 1 {
		fmt.Printf("hello world\t")
		for x, y := range os.Args {
			if x == 0 {
				continue
			}
			fmt.Printf("%v\t", y)
		}

	} else {
		os.Exit(-1)
	}
}
