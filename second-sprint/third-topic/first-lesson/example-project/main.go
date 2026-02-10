package main

import (
	"example-project/mathSlice"
	"fmt"
)

func main() {
	slice := mathSlice.Slice{1, 2, 3}
	fmt.Println(slice)
	fmt.Println("Slice's sum: ", mathSlice.SumSlice(slice))

	mathSlice.MapSlice(slice, func(i mathSlice.Element) mathSlice.Element {
		return i * 2
	})
	fmt.Println("Slice multiplied by 2: ", slice)
	fmt.Println("New slice's sum: ", mathSlice.SumSlice(slice))

	fmt.Println("Slice convolution by multiplying: ",
		mathSlice.FoldSlice(slice,
			func(x mathSlice.Element, y mathSlice.Element) mathSlice.Element {
				return x * y
			}, 1))

	fmt.Println("Convolution of a slice by addition: ",
		mathSlice.FoldSlice(slice,
			func(x mathSlice.Element, y mathSlice.Element) mathSlice.Element {
				return x + y
			}, 0))
}
