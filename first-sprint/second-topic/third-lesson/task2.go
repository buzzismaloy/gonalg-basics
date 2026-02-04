package main

import "fmt"

const (
	i         = 10
	f         = 1.5
	i64 int64 = 88
)

func main() {
	var v = 45
	var iplusf = i + f
	var ii64 = i + i64
	var fi64 = f + i64
	var i64v = i64 + v
	var iv = i + v
	var fv = f + v

	fmt.Println("i + f =", iplusf, "i + i64 =", ii64, "f + i64 =", fi64, "i64 + v =", i64v, "i + v =", iv, "f + v =", fv)
}
