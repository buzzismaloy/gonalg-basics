package main

import (
	"fmt"
	"math"
)

type figures int

const (
	square figures = iota
	circle
	triangle
)

func area(figure figures) (func(float64) float64, bool) {
	switch figure {
	case square:
		return func(x float64) float64 {
			return x * x
		}, true
	case circle:
		return func(x float64) float64 {
			return x * x * math.Pi
		}, true
	case triangle:
		return func(x float64) float64 {
			return (x * x * math.Sqrt(3)) / 4
		}, true
	default:
		return nil, false
	}
}

func main() {
	myFigure := square
	ar, ok := area(myFigure)
	if !ok {
		fmt.Println("Error")
		return
	}

	myArea := ar(2)
	fmt.Println(myArea)
}
