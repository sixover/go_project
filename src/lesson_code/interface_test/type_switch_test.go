package interface_test

import (
	"fmt"
	"testing"
)

func JudgeValueType(v interface{}) string {
	switch x := v.(type) {
	case int:
		return fmt.Sprintf("the type is %T\n", x)
	case string:
		return fmt.Sprintf("the type is %T\n", x)
	case rune:
		return fmt.Sprintf("the type is %T\n", x)
	case bool:
		return fmt.Sprintf("the type is %T\n", x)
	case byte:
		return fmt.Sprintf("the type is %T\n", x)
	default:
		return fmt.Sprintf("unknown type")
	}
}

func TestTypeSwitch(t *testing.T) {
	t.Log(JudgeValueType(1))
	t.Log(JudgeValueType("aaaa"))
}
