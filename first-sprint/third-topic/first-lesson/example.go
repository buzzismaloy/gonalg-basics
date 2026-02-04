package main

import "fmt"

func main() {
	a := 0.10000001
	// initialisation and main condition
	if b := float32(a); b > float32(0.1) {
		fmt.Println("variable a is greater than float32(0.1)")
	}

}
