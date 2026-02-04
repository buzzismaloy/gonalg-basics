package main

import "fmt"

const (
	Black = "black"
	Gray  = "gray"
	White = "white"
)

const (
	First  = 0
	Second = 1
	Third  = 2
)

const (
	FirstItem  = 0
	SecondItem = 0
)

func main() {
	fmt.Println("Black != Gray:", Black != Gray)
	fmt.Println("First != Second:", First != Second)
	fmt.Println("FirstItem != SecondItem:", FirstItem != SecondItem)
}
