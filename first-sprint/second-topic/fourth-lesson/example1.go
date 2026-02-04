package main

import "fmt"

const program = "My app"

var name string
var version = "v1.0.0"

func main() {
	if len(name) == 0 {
		fmt.Println("The global variable name is empty")
	}
	name = "Vasya"
	fmt.Println("name =", name)
	fmt.Println("Welcome to ", program, version)
}
