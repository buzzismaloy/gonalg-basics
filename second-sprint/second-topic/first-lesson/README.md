# First lesson

## Functions

Principle 4 of the structured programming methodology states: "Repetitive program fragments can be arranged in the form of subprograms (procedures and functions)."

A **function** is a logically complete piece of code with one input and one output in the control flow. This section can be used repeatedly by referring to it by name.

### Function declaration

A function declaration is often called a signature.

Syntax:

```go
func MyFunction(ar1 arg1type, arg2 arg2type, ..., argN argNtype) resultType {

}
```

The result of the function can also be named:

```go
func DivideX(x int) (half int) {
    half = x / 2
    return // then you don't have to specify the name in the return statement.
}
```

Function example:

```go
func Cube(x int) int {
    return x * x * x
}

// usage
result := Cube(5)
```

### Parameters

***Go has a special syntax for functions that can be called with a variable number of arguments (**variadic functions**). The parameter that accepts such arguments should be placed last in the list, and an ellipsis should be placed before its type.***

```go
func Sum(x ...int) int
```

Inside the function, this parameter is treated as a numbered sequence of arguments (slice).

```go
func Sum(x ...int) (res int) {
    for _, v := range x {
        res += v
    }
    return
}

// usage
sum := Sum(2, 3, 5, 7, 1, 3)
```

If you call this function without `Sum()` arguments, the parameter `x` will take the value `nil`. Then the loop will not go through any iterations, and the function will return `0`.

### Return values

The function can return not one, but several values of different types.

```go
func foo() int, string, bool, int, error {}
```

When calling such a function, variables must be provided to which all these values must be assigned.

```go
x, str, flag, y, err := foo()
```

And if some values are not needed, you can use the _ variable.

```go
x, str, _, y, _ := foo()
```

Here is a function that finds the index of a letter in a string and returns false with the second argument if the letter is not found:

```go
func Index(st string, a rune) (index int, ok bool) {
    for i, c := range st {
        if c == a {
            return i, true
        }
    }
    return // default values returned
}
```

If the number and type of values returned by a function exactly match the parameters of another function,

```go
func foo() (int, int) {}
func bar(x int, y int) {}
```

then this call syntax is allowed:

```
bar(foo())
```

### Recursive functions

Here is a textbook example of recursively calculating n!, the factorial of a number:

```go
func fact(n int) int {
    if n == 0 {
        return 1
    } else {
        return n * fact(n-1)
    }
}
```

And here are the Fibonacci numbers:

```go
func Fib(n int) int {
    switch {
    case n <= 1:
        return n
    default:
        return Fib(n-1) + Fib(n-2)
    }
}
```

It should be remembered that in Go, calling a function has a certain computational cost, as well as memory costs, because at least you need to copy the arguments. Therefore, a lot of nested function calls can lead to a decrease in program performance and memory overflow.

Iterative algorithms will work faster. For comparison, we present an iterative implementation (based on loops) of the above examples:

```go
func fact(n int) int {
    res := 1
    for n > 0 {
        res *= n
        n--
    }
    return res
}

func Fib(n int) int {
    a, b := 0, 1
    for n > 0 {
        a, b = b, a+b
        n--
    }
    return a
}
```

However, this does not mean that recursive algorithms are inapplicable. In some cases, they can be more useful, simpler, and make the code clearer.

The example of working with recursive traversal of all files in a given directory, and the directory may contain nested subdirectories is given in `example.go`.

### First class function

Functions in Go are in no way inferior to other classes of objects. A function has a type and a value. A function can be assigned to a variable, or passed as an argument to another function. A function can return another function as a value.
The type of a function is visible in its signature, that is, it is defined as a set of types and the number of arguments and return values.

For example, this function:

```go
func Say(animal string) (v string) {
    switch animal {
    default:
        v = "heh"
    case "dog":
        v = "gav"
    case "cat":
        v = "myau"
    case "cow":
        v = "mu"
    }
    return
}
```

has type:

```go
func(string) string

// it can be assgined to variable with the same type
var voice func(string)string
voice = Say
```

You can write a higher-order function with a parameter of this type:

```go
func Print(who string, how func(string) string) {
    fmt.Println(how(who))
}

// usage
Print("dog", Say)
```

There is a literal syntax for the function. The function can be created locally, without declaring or naming it in the declaration block:

```go
f := func(s string) string { return s }

// You can even use a literal as an argument when calling:
Print("dog", func(s string) string { return s })
```

This is what is also called an anonymous or lambda function.

You can write a function that returns functions with values:

```go
func Do(say bool) func(string) string {
    if say {
        return Say
    }

    return func(s string) string { return s }
}

// usage
Print("dog", Do(true))
```

### Closures

Go is a **lexically scoped language**. This means that variables defined in the surrounding visibility blocks (for example, global variables) are always available to the function, and not just for the duration of the call. We can assume that the function remembers them.
The lexical scope and anonymous functions allow you to implement **closures**.

Here is a classic example of an iterator of even numbers built on closure:

```go
func Generate(seed int) func() {
    return func() {
        fmt.Println(seed) // closure gets outer variable seed
        seed += 2 // variable is modified
    }
}

func main() {
    iterator := Generate(0)
    iterator()
    iterator()
    iterator()
    iterator()
    iterator()
}
```

A closure binds an external variable to itself. After exiting the external Generate function, it is not destroyed, but remains bound to the closure function, and its value is preserved between function calls.

output:

```
0
2
4
6
8
```

Fibonacci numbers with closures:

```go
func fib() int {
    x1, x2 := 0, 1

    return func() int {
        x1, x2 = x2, x1+x2
        return x1
    }
}

func main() {
    f := fib() // got function-closure. f() - got x1, x2. x1 = 0, x2 = 1
    fmt.Println(f()) // x1 = 1, x2 = 1
    fmt.Println(f()) // x1 = 1, x2 = 2
    fmt.Println(f()) // x1 = 2, x2 = 3
    fmt.Println(f()) // x1 = 3, x2 = 5
    fmt.Println(f()) // x1 = 5, x2 = 8
    fmt.Println(f()) // x1 = 8, x2 = 13
}
```

output:

```
1
1
2
3
5
8
13
```

Such functions are sometimes called **generators**. They give out a new value of a sequence each time they are called.
Closures are quite useful. They allow you to implement certain design patterns simply and gracefully. Nevertheless, in order to use closures effectively, you need to understand how they work.
Here is a more practical example of using a closure. Let's create two wrapper functions, one of which will count the number of calls, and the second — the execution time of the function. This example is given in `example1.go`.

Here is another interesting example of using closure. Let's recall the `PrintAllFilesWithFilter` function which is given in `task.go`, which we recently worked with. Its disadvantage is that the `filter` parameter is passed with each recursive call. You can get rid of this by using an anonymous function as a closure. This example is given in `example-task.go`. Now the unchanging parameter is not copied at each step of the recursion. The closure simply refers to its value, increasing the speed of the program and reducing the likelihood of error.

**This approach is often used in Go web development, when a group of handler functions are chained together, sharing responsibility for certain actions.***

### Special functions

The entry point to the program is the `main()` function. It must exist in a single form and in any executable program on Go. `main()` does not accept arguments and does not return values.
Go has built-in functions, for example: `make(), new(), len(), cap(), delete(), close(), append(), copy(), panic(), recover()`. These are not library functions. They don't fully follow the rules for user functions. They may not have a signature, and their use is documented in the language specification, which is the fundamental document for Go.
This function is also described in the basic syntax of the language:

```go
func init() {...}
```

Several such functions can be declared in a package or even in a single file. They will be called once during package initialization, after assigning global variables, in the order in which they are provided to the compiler (found in the source text). There is no direct call to the `init()` function in the program code.
These functions are used to create the environment necessary for the package to work correctly.

Here is a simple example:

```go
var name, surname string

func init() {
    name = "John"
}
func init() {
    if surname == "" {
        surname = "Doe"
    }
}
func main() {
    fmt.Println("Hello " + name + " " + surname)
}
```
