# Second lesson

## FizzBuzz

FizzBuzz.go file shows the realisation of the simple program FizzBuzz how the ***clean*** is Go code.

## OOP

The OOP realisation has the so-called "duck typing": "If something swims like a duck, quacks like a duck, and flies like a duck, then it's most likely a duck." It is enough to implement a set of methods for a type so that it starts automatically satisfying all interfaces with similar method signatures:

```go
type Stringer interface {
    String() string
}

type myType int

// myType implements interface Stringer
func (t myType) String() string {
    // implementation of type myType as a string
}
```

## Exceptions

Go uses a different approach that brings more clarity to the error handling process. The fact is that functions in Go can return more than one value. This property is actively used by developers, using the `error` interface as the last return value.

The example of how golang handles errors is presented in `exceptions.go`

## Panic

Go also so-called mechanism as `Panic`
If the construction above is a typical way to check the performance of a function, then panic is thrown only when the executing code gets into a non—standard situation that cannot be handled.
Go provides the ability to intercept and handle panics. This uses the `defer` design and the built-in `recover` function.

`defer` is another unusual language concept that executes blocks of code when exiting a function, for example, to close files upon completion of work with them. `defer` can be considered as a replacement for destructors/context managers in other languages (`try_with_resources` from Java, `with` from Python). `defer` is performed even in case of panic, when an emergency shutdown of functions occurs.

```go
func foo() {
    // paniic
    panic("unexpected!")
}
//...
    // executes after panic
    defer func() {
        if r := recover(); r != nil {
            // panic handling, in the variable r will be string "unexpected"
        }
    }()
    // inside the foo panic is triggered
    foo()
```

It may seem that panic is very similar to the exception mechanism, but it is not. By throwing an `exception`, the function usually waits for the exception to be caught by the exception handler above. However, panic is not always intercepted. If `recover()` is not called in case of panic, when returning control back down the function stack, the program will terminate with an error.

Let's determine what happens if in the previous example (which is presented in the `exceptions.go`) function `div` wont check division by 0.

```go
func div(a, b int) int {
    return a / b
}
```

Calling `div(10, 0)` will terminate the program with the error `panic: runtime error: integer divide by zero`. This error indicates that a panic has occurred during operation. If error checking is a typical way to check the performance of a function, then panic is thrown only when the executing code gets into an unusual situation that cannot be handled.

```go
func main() {
    var a []int

    a[0] = 7 // panic is triggered here
}
```

In the example above panic is triggered because array is empty:

```
panic: runtime error: index out of range [0] with length 0

goroutine 1 [running]:
main.main()
    /tmp/sandbox1542797293/prog.go:10 +0x18

Program exited.
```

You can create an emergency yourself by calling the `panic` function with any type of parameter. By default, panic will go up the stack and terminate all functions until it completes the `main` function, and with it the entire process. The use of the `panic()` function should be treated with caution, as the postulate **"Do not panic"** reminds. There is no need to use panic where you can simply return an error and then process it.

`recover` is a function that allows you to restore program execution in case of a panic. If an emergency has occurred at the time of the `recover` call, `recover` terminates it and returns the error value (the argument when calling panic). If there was no emergency, `recover` does nothing and returns `nil`.
Example of `recover` realisation is presented in `recover-example.go`

## Tools for testing

The `testing` library was mentioned above. In Go, it is customary to place files with unit tests directly in the package whose functions you are testing. For example, if the code you want to cover with tests is located in the `foo.go` file, you need to create the `foo_test.go` file for the tests:

```bash
foo/ # foo package
    foo.go # file with tested code
    foo_test.go # file with tests
```

`foo.go` contents:

```go
package foo

func Foo() string {
    return "bar"
}
```

In `foo_test.go` we implement functions of certain signature:

```go
package foo

import (
    "testing"
)

func TestFooFunc(t *testing.T) {
    expectedFooResult := "bar"
    if actualFooResult := Foo(); actualFooResult != expectedFooResult {
        t.Errorf("expected %s; got: %s", expectedFooResult, actualFooResult)
    }
}
```

To execute tests just run `go test`:

```bash
PASS
ok      github.com/Yandex-Practicum/go-freetrack/00_intro/testing/foo   0.240s
```

## Multithreading in Go

As mentioned earlier, multithreading in Go is implemented according to the CSP (Communicating Sequential Processes) model. With this approach, the program is a set of simultaneously running subtasks that communicate through communication channels. The tasks in Go are goroutines, communication is organized through channels.

A **goroutine** is a lightweight thread that takes up much less memory than an OS thread. The Go runtime can execute multiple goroutines on the same operating system thread and quickly switch from executing one goroutine to another due to their small size. The preemptive scheduler tries to evenly distribute processor time between the goroutines.

**Channels** are the second key element in multithreading on Go. Channels not only enable streams to exchange data, but also serve to synchronize their operation. One goroutine can write data to the channel, and the other goroutine can read it. In addition, the standard library has additional thread synchronization primitives.


| CSP | Go |
|----------|----------|
| subtask | goroutine |
| communication channel | channel |
