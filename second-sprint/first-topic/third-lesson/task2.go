package main

import "fmt"

func RemoveDuplicates(input []string) []string {
	m := make(map[string]int, len(input))

	for i, value := range input {
		if idx, ok := m[value]; ok {
			if len(input) != 0 && idx < len(input) {
				input = append(input[:idx], input[idx+1:]...)
			}
		}
		m[value] = i
	}

	return input
}

func main() {
	input := []string{
		"cat",
		"dog",
		"bird",
		"dog",
		"parrot",
		"cat",
	}
	fmt.Println("Before removing:", input)
	input = RemoveDuplicates(input)
	fmt.Println("After removing", input)
}
