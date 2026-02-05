package main

import "fmt"

func main() {
	sentence := "Πολύ λίγα πράγματα συμβαίνουν στο σωστό χρόνο, και τα υπόλοιπα δεν συμβαίνουν καθόλου"
	freq := make(map[rune]int)

	for _, value := range sentence {
		freq[value]++
	}

	for key, value := range freq {
		fmt.Printf("Character %c is found %d times\n", key, value)
	}
}
