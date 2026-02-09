package main

import "fmt"

var Global = 5

// second solution
func PrintGlobal() {
	defer func(number int) {
		Global = number
	}(Global)
	Global = 31
	fmt.Println(Global)
}

func main() {

	// first solution
	/*defer func() {
		Global = 7
		fmt.Println(Global)
		Global = 5
	}()*/

	fmt.Println(Global)
	PrintGlobal()
	fmt.Println(Global)

}
