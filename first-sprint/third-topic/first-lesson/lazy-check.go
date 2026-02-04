package main

import "fmt"

func main() {
	a, b := 1, 0

	incB := func() bool {
		b = b + 1
		return true
	}

	if a == 1 || incB() {
		fmt.Println("Hello")
	}

	fmt.Println("a =", a, "b =", b)
}
