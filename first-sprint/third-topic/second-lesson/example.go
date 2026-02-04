package main

import "fmt"

func main() {
	group := 0

	for i := 0; i < 20; i++ {
		switch {
		case i%2 == 0:
			if i%10 == 0 {
				group++
				break
			}
			fmt.Printf("%02d: %d\n", group, i)
		default:
		}
	}

}
