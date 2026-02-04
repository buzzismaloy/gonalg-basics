package main

import "fmt"

const (
	_ = iota * 10 // обратите внимание, что можно пропускать константы
	ten
	twenty
	thirty
)

const (
	hello = "Hello, world!" // iota равна 0
	one   = 1               // iota равна 1

	black = iota // iota равна 2
	gray
)

func main() {
	fmt.Println(ten, twenty, thirty)
	fmt.Println(one, black, gray)
}
