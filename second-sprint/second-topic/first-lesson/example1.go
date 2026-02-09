package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

func CountCall(f func(string)) func(string) {
	counter := 0
	funcname := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()

	return func(s string) {
		counter++
		fmt.Printf("Function %s called %d times\n", funcname, counter)
		f(s)
	}
}

func metricTimeCall(f func(string)) func(string) {
	return func(s string) {
		start := time.Now()
		f(s)
		fmt.Println("Execution time:", time.Now().Sub(start))
	}
}

func myprint(s string) {
	fmt.Println(s)
}

func main() {
	counterPrinted := CountCall(myprint)
	counterPrinted("Hi")
	counterPrinted("Hello there")

	countAndMetricPrint := metricTimeCall(counterPrinted)
	countAndMetricPrint("Hello")
	countAndMetricPrint("there")
}
