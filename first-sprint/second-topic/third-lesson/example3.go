package main

import "fmt"

const (
	Black = iota
	Gray
	White
)

// the counter is reset to zero

const (
	Yellow = iota
	Red
	Green = iota // this assignment will not reset iota.
	Blue
	Purple
)

func main() {
	fmt.Println(Black, Gray, White)
	fmt.Println(Yellow, Red, Green, Blue, Purple)
}
