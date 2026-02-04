package main

import (
	"fmt"
	"time"
)

func main() {
	date := time.Date(2013, 12, 1, 0, 0, 0, 0, time.UTC)

	switch y := date.Year(); {
	case y >= 1946 && y < 1965:
		fmt.Println("Hi boomer!")
	case y >= 1965 && y < 1981:
		fmt.Println("Hi X-gen!")
	case y >= 1981 && y < 1997:
		fmt.Println("Hi, millennial!")
	case y >= 1997 && y < 2013:
		fmt.Println("Hi zoomer!")
	case y >= 2013:
		fmt.Println("Hi alpha!")
	default:
		fmt.Println("Hi!")
	}
}
