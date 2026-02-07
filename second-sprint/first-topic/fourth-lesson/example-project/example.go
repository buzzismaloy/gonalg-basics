package main

import (
	"example-project/foo"
	"fmt"
)

func main() {
	//f := foo.privateFoo {} // compilation error
	f := foo.NewPrivateFoo()
	fmt.Println(f.Value) // field Value is exported so you can use it
}
