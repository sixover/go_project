package switch_test

import (
	"fmt"
	"testing"
)

func TestSwitchMultiCaseFunction(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 1, 3:
			fmt.Println("its odd")
		case 2, 4:
			fmt.Println("its even")
		default:
			fmt.Println("unknown branch!")
		}
	}
}

func TestSwitchUseForIf(t *testing.T) {
	for i := 0; i < 6; i++ {
		switch {
		case 0 <= i && i < 2:
			fmt.Println("its between 0 and 2")
		case 2 <= i && i < 5:
			fmt.Println("its between 2 and 5")
		default:
			fmt.Println(i)
		}
	}
}
