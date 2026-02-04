package main

import "fmt"

func foo() {
	panic("unexpected") // panic is triggered
}

func main() {
	// executes at the end of main lifetime
	defer func() {
		// call recover and compare the result with nil
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	foo() // inside the foo panic is triggered
	fmt.Println("You won't see this message as panic is triggered in foo")
}
