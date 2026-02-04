package main

import "fmt"

const (
	one    = 2*iota + 1 // iota = 0 + 1
	three               // iota = 1 + 2
	five                // iota = 2 + 3
	seven               // iota = 3 + 4
	nine                // iota = 4 + 5
	eleven              // iota = 5 + 6
)

func main() {
	fmt.Println(one, three, five, seven, nine, eleven)
}
