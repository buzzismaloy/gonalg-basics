package main

import "fmt"

func main() {
	i := 10

	if i == 10 {
		i += 5
		if i == 15 {
			i := 7
			fmt.Println("i in nested brackets equals to", i)
		}
	}
	fmt.Println("i outside the brackets in main equals to", i)
}
