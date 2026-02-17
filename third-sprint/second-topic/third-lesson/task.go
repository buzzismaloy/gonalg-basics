package main

import (
	"fmt"
	"strings"
)

func Mul(a any, b int) any {
	switch v := a.(type) {
	case int:
		return v * b
	case string:
		/*newString := strings.Builder{}
		for i := 0; i < b; i++ {
			newString.WriteString(v)
		}
		return newString*/
		return strings.Repeat(v, b)
	case fmt.Stringer:
		return strings.Repeat(v.String(), b)
	default:
		return nil
	}
}

func main() {
	result1 := Mul(2, 3)
	result2 := Mul("ok", 4)

	fmt.Printf("result1: %d\nresult2: %s", result1, result2)
}
