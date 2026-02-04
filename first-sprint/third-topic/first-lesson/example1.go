package main

import "fmt"

func main() {
	a := -100
	switch {
	case a > 0:
		if a%2 == 0 {
			break
		}
		fmt.Println("Odd positive value received")

	case a < 0:
		fmt.Println("Nevative value received")
		fallthrough
	default:
		fmt.Println("Default value handling")
	}
}
