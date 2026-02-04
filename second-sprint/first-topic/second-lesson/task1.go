package main

import "fmt"

func main() {
	var size int = 100
	var num int = 10

	slice := make([]int, size)

	for i := 0; i < size; i++ {
		slice[i] = i + 1
	}

	slice = append(slice[:num], slice[size-num:]...)
	size = len(slice)

	for i := range slice[:size/2] {
		slice[i], slice[size-i-1] = slice[size-i-1], slice[i]
	}

	fmt.Println(slice, size, cap(slice))
}
