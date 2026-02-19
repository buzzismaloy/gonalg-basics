# Second lesson

## Panic and recover functions

Now you know for sure that errors can occur during program execution.

- Sometimes errors aren't a big deal—for example, the user passed invalid data. In this case, simply logging them is sufficient.
- Sometimes errors can cause problems when information cannot be retrieved from external resources. In this case, you can notify the user or try retrieving the information again.
- Sometimes errors can be so severe that continuing the program is impractical or even dangerous. For example, a database connection failed, or errors occurred in configuration files.

The latter situations are considered abnormal and require immediate program termination. This is accomplished using the panic mechanism. You're already familiar with it: a program panics, for example, when dereferencing a `nil` pointer or performing a failed type cast. In this chapter, we'll look at panic in more detail.

In an abnormal situation, the program stops, deferred `defer` functions are called, and an error message is displayed. This message, in addition to the text, also indicates the state of the function call stack.

To create an emergency situation, call the built-in `panic` function. You can pass any type of value as an argument to `panic`:

```go
import (
    "fmt"
    "os"
)

func myfunc() {
    if _, err := os.ReadFile(`test.txt`); err != nil {
        panic(err)
    }
}

func main() {
    fmt.Println("start")
    myfunc()
    fmt.Println("finish")
}
```

output:

```
Старт
panic: open test.txt: no such file or directory

goroutine 1 [running]:
main.myfunc(...)
    /tmp/sandbox910916036/prog.go:10
main.main()
    /tmp/sandbox910916036/prog.go:16 +0x11e
```

The program didn't print the "finish" line. You can immediately see not only the cause of the error but also the line number where it occurred.
In addition to the `panic` function, the Go runtime can generate crashes for scenarios such as division by zero, null pointer dereference, and out-of-bounds access to an array or slice. Such errors can occur in any function. For example, it's inefficient to check for array out-of-bounds or for a `nil` pointer during every operation:

```go
import "fmt"

func mypanic() {
    var slice []string
    fmt.Println(slice[0])
}

func main() {
    mypanic()
}
```

If you run this fragment, the program will exit with the following error text:

```
panic: runtime error: index out of range [0] with length 0

goroutine 1 [running]:
main.mypanic()
    /tmp/sandbox852111250/prog.go:7 +0x18
main.main()
    /tmp/sandbox852111250/prog.go:11 +0x25
```

The panic occurred on line 7, in the function called on line 11.

***Using the `panic` function to handle all errors is not recommended. It's best to avoid `panic` in libraries. If such a function is necessary, it should be documented.***

`panic` is a resource-intensive mechanism that can't always be stopped. For this reason, panic is not recommended.

Here's an example description for the `MustCompile` function in the standard `regexp` package:

```go
func MustCompile(str string) *Regexp
// ...
MustCompile is similar to Compile but panics if the expression cannot be parsed. It simplifies the safe initialization of global variables holding compiled regular expressions.
```

Creating an emergency situation is possible when it's unclear how the program will proceed and it's better to terminate the process now than to risk more serious problems. For example, panic is appropriate in tests and scenarios when the program is unable to read the configuration file at startup. `Panic` is also often used in `switch-case` statements if the variable value doesn't match any of the options.

```go
switch srvType {
case LocalHost:
    // ...
case RemoteHost:
    // ...
default:
    panic(fmt.Sprintf(`Unknown type of the server %d`, srvType))
}
```

It is recommended to log each panic and enable the developer alert system for quick notification of the problem.

How does a panic differ from a normal program termination? Go has another function, `recover`, which allows you to resume program execution in the event of a panic. If a panic occurs when you call `recover`, `recover` terminates it and returns the error value (the argument when calling `panic`). If there was no panic, `recover` does nothing and returns `nil`.

***Note***: Don't pass `nil` when calling `panic` — if `recover` returns it, you won't know there was a panic.

Where should you call `recover` if the program has stopped executing? Since `defer` functions are called when a panic occurs, we'll insert the call to `recover` there.
As a reminder, `defer` calls will be issued regardless of whether a panic occurs. We discussed this in detail in the section on `defer`:

```go
import "fmt"

func mypanic() {
    defer func() {
        if p := recover(); p != nil {
            fmt.Println(`Panic occurred: `, p)
        }
    }()
    panic(`emergency situation`)
}

func main() {
    fmt.Println("Start")
    mypanic()
    // main func will continue its work as the recover function was used
    fmt.Println("Finish")
}
```

output:

```
Start
Panic occurred: emergency situation
Finish
```

The combination of `panic` and `recover` functions is similar to the `try-catch` exception mechanism in other languages. However, while standard error handling requires no additional resources, panicking results in stack unwinding, which is an expensive operation.
Although the `recover` function allows program execution to continue, it should only be used in justified cases. For example, a standard web server uses `recover` to prevent a panic in one handler from terminating the entire process.

It's also worth remembering about possible memory leaks and undefined global variables after exiting a crash, which can lead to new conflicts. There are situations where panics cannot be caught, and the program will definitely terminate. These include, for example, situations related to concurrency or running out of memory.

## Key Points

- Panic interrupts normal program execution and is therefore only applicable in emergency situations.
- Panic causes unexpected resource overhead associated with stack unwinding.
- Deferred `defer` calls are always called. This is where panics should be caught with the `recover` function.
- Not every panic can be recovered.
- Even if a panic must be triggered within a package, it should not be released outside the package.
