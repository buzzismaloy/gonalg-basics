package main

import "fmt"

func main() {
	input := "The quick brown 󰌌 jumped over the lazy 󰌌"
	n := 0

	runes := make([]rune, len(input))

	for _, r := range input {
		runes[n] = r
		n++
	}

	runes = runes[0:n]

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-i-1] = runes[n-i-1], runes[i]
	}

	output := string(runes)

	fmt.Println(output)
}
