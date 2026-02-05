package main

import "fmt"

func firstSolution(number int, array []int) []string {
	var returnSlice []string

	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i]+array[j] == number {
				msg := fmt.Sprintf("%d + %d, %d : %d", array[i], array[j], i, j)
				returnSlice = append(returnSlice, msg)
			}
		}
	}

	return returnSlice
}

func secondSolution(number int, array []int) []string {
	var returnSlice []string
	m := make(map[int][]int)

	for i, value := range array {
		need := number - value

		if indices, ok := m[need]; ok {
			for _, idx := range indices {
				msg := fmt.Sprintf("%d + %d, %d : %d", array[idx], value, idx, i)
				returnSlice = append(returnSlice, msg)
			}
		}

		m[value] = append(m[value], i)
	}

	return returnSlice
}

func main() {
	var k int = 8
	A := []int{1, 2, 4, 5, 6, 10, 5, 2, 3, 7, 0, 8, -2}

	fmt.Println(A)

	firstSlice := firstSolution(k, A)
	if firstSlice != nil {
		fmt.Println("The first solution:\n")
		for _, msg := range firstSlice {
			fmt.Println(msg)
		}
	} else {
		fmt.Println("Solution hasn't been found")
	}

	secondSlice := secondSolution(k, A)
	if secondSlice != nil {
		fmt.Println("\n\nThe second solution:\n")
		for _, msg := range secondSlice {
			fmt.Println(msg)
		}
	} else {
		fmt.Println("Solution hasn't been found")
	}
}
