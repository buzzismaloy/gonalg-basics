package main

import (
	"bufio"
	"fmt"
	"os"
)

func f(counter *int) {
	*counter++
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Interaction counter")

	cnt := 0
	for {
		fmt.Print("-> ")
		_, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		f(&cnt)

		fmt.Printf("User input %d lines\n", cnt)
	}
}
