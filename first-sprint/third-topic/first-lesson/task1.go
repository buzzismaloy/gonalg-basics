package main

import "fmt"

func main() {
	age := 1950

	switch {
	case age >= 1946 && age <= 1964:
		fmt.Println("Hi, boomer!")

	case age >= 1965 && age <= 1980:
		fmt.Println("Hi, X!")

	case age >= 1981 && age <= 1996:
		fmt.Println("Hi, millennial!")

	case age >= 1997 && age <= 2012:
		fmt.Println("Hi, zoomer!")

	case age >= 2013:
		fmt.Println("Hi, alpha!")

	default:
		fmt.Println("default handling")
	}
}
