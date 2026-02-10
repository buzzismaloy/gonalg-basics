package main

import (
	"fmt"

	"example/toppackage/middlepackage/bottompackage/mathxxx"
)

func main() {
	if sum := mathxxx.AddInts(3, 2); sum != 5 {
		panic(fmt.Sprintf("sum must be equal 5; got %d", sum))
	}

	fmt.Println("Well done!")
}
