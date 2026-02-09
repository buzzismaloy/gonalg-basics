# Second lesson

## Deferred Call Operator

There is an operator in Go that allows you to schedule a deferred call, this is the `defer` instruction.
It was briefly reviewed earlier.

```go
resource := System.Acquire("resourceID")
defer System.Close(resource)
```

The `defer` operator is used inside functions, and its operand is a function call expression. We will call these functions "the deferred" and "deferred" to avoid confusion.
The `defer` instruction computes the arguments for the call, but the call does not. The call is executed immediately before the function that postponed it returns control.

```go
func EvaluationOrder(){
    defer fmt.Println("deferred")
    fmt.Println("evaluated")
}
```

output:

```
evaluated
deferred
```

The operator's work can be described as follows.

    1. The program is running normally.
    2. It's time for the defer statement to be executed.
    3. The operands of the deferred function are evaluated, if any.
    4. The function call, along with the values, is postponed to a special stack.
    5. The function continues to be executed. If the defer operator occurs, then repeat points 3 and 4.
    6. If the return operator occurs, the function evaluates its operands and stores the value in the buffer.
    7. If the stack of deferred calls is not empty, then we extract the function call from it and execute it.
    8. Repeat step 7 until the stack is empty.
    9. We exit the function by returning the value from the buffer.

It is important to understand that the result of the function is calculated before executing deferred calls.
There may be several pending calls. Then they are executed in reverse order, that is, starting from the one that was postponed last, since the calls were stacked.

```go
fmt.Println("Hello")
for i := 1; i <= 3; i++ {
    defer fmt.Println(i)
}
fmt.Println("World")
```

output

```
Hello
World
3
2
1
```

A deferred function may return a value that is not being used. Indeed, there is simply nowhere to return it.

```go
func VeryImportantFunc(s string, x, y, z int) (a int, b bool) {
    ...
}

// Go allows you to write:
defer VeryImportanttFunc(s,x,y,z)
```

A deferred function can also be anonymous and set literally. Recall that an anonymous function is a function specified by a literal at the place of use. In this case, the anonymous function is set immediately along with the call.
In this case, the variables of the deferred function may be available to it. A **closure** will occur. For example, if a deferred function has a named return value, the deferred function can change it.

```go
func unintuitive() (value string){
    defer func() {value = "На самом деле"}() // () - means the function is called
    return "Казалось бы"
}
```

Note that this only works with named return values. The following code outputs `"Казалось бы"`:

```go
func intuitive() (string){
    value := "Казалось бы"
    defer func() {value = "На самом деле"}()
    return value
}
```



What's the difference? In the first case, the function returns the `value` variable. When calculating the `return` operand, it is indeed assigned the value `"Казалось бы"`, but this variable is captured by the closure and changed in it. After that, it returns from the function.
In the second case, we have some hidden variable `ret`1, into which, when calling the `return` operator, the value of its operand is copied. After that, any actions with `value` will no longer be important.

It is also a common mistake to assume that the operands of a deferred function will be evaluated during its execution. This is not the case, they are calculated when executing the `defer` statement:

```go
package main

import "fmt"

func main() {
    a := "some text"
    defer func(s string) {
        fmt.Println(s)
    }(a)
    a = "another text"
}

// The program will print "some text".
```

The defer operator is most often seen with the paired functions Open()/Close(), Lock()/Unlock(). It is installed immediately after the capture of the resource, so as not to forget exactly.
Here is a classic example:

```go
// open file
file, err := os.OpenFile("file.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
if err != nil {
    log.Fatal(err)
}
// do not forget to close the file
defer file.Close()
// работаем с файлом
_, err = file.WriteString("")
if err != nil {
    log.Fatal(err)
}
```

**Example**:

Based on the `defer` function, we will measure the execution time of the function.
First, let's create a function that will measure the execution time and display it on the screen.

```go
    func metricTime (start time.Time) {
        fmt.Println(time.Now().Sub(start))
    }

// usage inside another function

    func VeryLongTimeFunction () {
        defer metricTime(time.Now())
        // some actions
    }
```

## defer and panic

Various unforeseen circumstances may occur during the execution of the program, which will make further operation of the function impossible. In this case, the execution of the function is immediately terminated, the panic is passed to the calling function and then up the stack until the execution of the program is completed.
However, this process changes if the panicked function has pending calls. They will be executed after exiting the function and may realize that a panic has occurred.
This is similar to the exception mechanism in C++ and Python, but it is highly discouraged to use `panic` for normal operation. Panic should be triggered only when the program execution really cannot continue and it must be completed.

```go
func PanicingFunc () {
    defer func(){
        if r := recover(); r != nil {  // recover stops panic and returns the reason of panic
            fmt.Println("Panic is caught", r)
        }
    }()

    panic("Мне здесь совсем ничего не нравится!")
    // panic() calls panic inside function
    // as an argument panic() takes the reason of the panic. It will be returned to recover()
}
```

Without using the defer operator, it would have been impossible to stop the panic. It allows you to insert yourself into the stack of function calls and stop it. Note that not all panic can be restored. Sometimes there are special situations when recover does not work.
