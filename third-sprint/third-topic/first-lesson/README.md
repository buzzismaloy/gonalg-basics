# First lesson

## Errors

In this topic, you'll learn about `error` types and how to:

- handle errors;
- define your own error types and generate errors when emergency situations occur;
- pack errors;
- create panic errors and recover from program failure.

Errors can occur when running any program. In some cases, this is due to invalid input data or inaccessible system resources. In others, it's due to poorly written code, which can lead to crashes and program termination.
Let's look at an example. Suppose a program asks the user for two numbers and divides one by the other. Obviously, it needs to check whether the divisor is zero. If such a check exists and the user enters zero, it's enough to ask them to enter another number. If there's no check, the program will terminate with a crash.
No programming language guarantees that errors won't occur, so any language has tools for handling them. For example, one approach is to use exceptions. You can catch errors anywhere in the program and decide what to do next.

***Go takes a different approach to error handling. A program detects an error as soon as it occurs and must handle it immediately. This requires additional code but reduces the number of crashes.***

### Error type

We've already mentioned the `error` type in passing. Indeed, the `if err != nil {}` construct is practically a hallmark of the language and appears in every Go program.
The `error` type is an interface type:

```go
type error interface {
    Error() string //  this method must return the text of the error.
}
```

Because functions can return multiple values, errors can easily be caught in them. Using `error` as the last return value of a function is a very common pattern in Go. If the `error` return value is `nil`, the function has completed gracefully. Otherwise, the remaining values of the function cannot be used, and the error must be handled.
Let's look at an example:

```go
if data, err := os.ReadFile(`nothing.txt`); err != nil {
    // the 'Error() string' method will be called which transfmorms the error to string
    fmt.Println(err)
} else {
    fmt.Println(string(data))
}
```

Go typically uses the "early exit" pattern, so the code above is better used like this:

```go
func ReadTextFile() (string,error) {
    if data, err := os.ReadFile(`nothing.txt`); err != nil {
        // the 'Error() string' method will be called which transfmorms the error to string
        fmt.Println(err)
        return "", err
    }
    fmt.Println(string(data))
    return string(data), nil
}
```

If an error occurs, there's likely no point in continuing, and it's best to exit as soon as possible. This also avoids cluttering the code with unnecessary `else` statements and curly braces, making it more idiomatic.

If you run this code with the missing file `nothing.txt`, the program will print `open nothing.txt: no such file or directory`. The return value `data` is not processed in any way.

The Go standard library includes the `errors` package for handling errors. To create a variable of type `error`, you need to call the `New` function, which takes a string as a parameter. For example, in the code above, we could create our own error instead of returning the one returned from the function.

```go
func ReadTextFile()  (string, error) {
    if data, err := os.ReadFile(`nothing.txt`); err != nil {
        // the 'Error() string' method will be called which transfmorms the error to string
        fmt.Println(err)
        return "", errors.New("some_file_process_func: read file error")
    }
    fmt.Println(string(data))
    return string(data), nil
}
```

However, this approach has its drawbacks. Multiple functions in a package can return the same errors. Furthermore, the error generation function will be called every time it's used, and the error variable will be re-created multiple times.

Therefore, static error generation, at module initialization, is most often used:

```go
// error is statically created
var ErrFileReading = errors.New("read_text_file: read file error") //  It is a good practice to start the error text with the name of the package where it is declared, this will make it easier to find

func ReadTextFile() (string, error) {
    data, err := os.ReadFile(`nothing.txt`)
    if err != nil {
        // the 'Error() string' method will be called which transfmorms the error to string
        fmt.Println(err)
        return "", ErrFileReading
    }
    fmt.Println(string(data))
    return string(data), nil
}
```

If you need to generate an error with additional information, you can use the `fmt.Errorf` function, which works like `fmt.Sprintf` but returns the error instead of a string. In Go, it's common practice to begin error text with a lowercase letter, as errors can be concatenated.

```go
func ReadTextFileByName(filename string) (string, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        // return error in russian
        return ``, fmt.Errorf(`Hе удалось прочитать файл (%s): %v`, filename, err)
    }
    return string(data), nil
}
```

### Comparison of errors

Almost any function can return various errors. It's helpful to be able to analyze them and understand what's happening for future reference. For example, a database might return a `NoRows error` if the requested data isn't found, but this isn't always a violation of the program's logic.
Errors can be compared just like other variables:

```go

data, err := ReadTextFile()
// check if the error is returned
if err != nil {
    if err == ErrFileReading {
        fmt.Println("unable read file")
        return
    }
    fmt.Println("unknown error")
    return
}
```

However, you can't compare an error to anything other than nil if it was generated dynamically.

Next, we'll learn how to resolve this issue.

### Error wrapping

In the previous section, we saw two approaches to creating errors: static and dynamic. Both approaches have drawbacks.
**Static** errors are quick to process, easy to read, and easy to compare, but they lack flexibility when additional information needs to be added to the error. **Dynamic** errors are created during program execution; they are slower and difficult to compare.
Starting with version 1.13, Go provides the ability to "wrap" errors. This means you can create a new error on top of an old one while still being able to restore the original error. This can be useful for creating your own error types based on existing ones. An example is provided below.

Recall the example of reading a configuration file, where one error is replaced by another. The original error text remains, but its type cannot be restored. To wrap the error, you need to use the `%w` specifier for the `Errorf` function.

Let's correct `%v to %w` in the `ReadTextFile` function:

```go
func ReadTextFile(filename string) (string, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        // return error in russian
        return "", fmt.Errorf(`не удалось прочитать файл (%s): %w`, filename, err)
    }
    return string(data), nil
}
```

In this case, you can restore the original error using the `errors.Unwrap` function and add additional checks.

Furthermore, the `github.com/pkg/errors` package contains the `Wrap and Wrapf` functions, which create a wrapped error.

Why wrap errors? Wrapping allows you to create a hierarchy of errors:

```go
// statically create error
var ErrFileReading = errors.New("read_text_file: read file error")

func ReadTextFileByName(filename string) (string, error) {
    if data, err := os.ReadFile(filename); err != nil {
        // the 'Error() string' method will be called which transfmorms the error to string
        fmt.Println(err)
        return errors.Wrapf(ErrFileReading, "file not exist %s", filename)
    }
    fmt.Println(string(data))
    return string(data), nil
}
```

Now, at the place where our function is called, let's compare:

```go
if errors.Is(err, ErrFileReading) {
    // do smt
}
```

The `Is` function compares errors, even wrapped ones! This allows us to handle wrapped errors and pass additional information in them.
It's considered good architectural practice for a package to return only its own errors. A package's set of errors is part of its contract for handling. By forwarding received errors from other packages, we increase package cohesion. Then, the error-handling side will have to handle not only the errors of our package, but also those of the packages it depends on.

### Custom Error Types

Sometimes, in addition to the error text, you need to pass additional information. Let's look at an example where we return the error's occurrence time along with the error text:

```go
import (
    "fmt"
    "time"
)
// Create a custom type that satisfies the error interface
// TimeError is a type for storing the time and error text.
type TimeError struct {
    Time time.Time
    Text string
}

// Error adds support for the error interface for the TimeError type.
func (te TimeError) Error() string {
    return fmt.Sprintf("%v: %v", te.Time.Format(`2006/01/02 15:04:05`), te.Text)
}

// NewTimeError returns a TimeError variable with the current time. func NewTimeError(text string) TimeError {
return TimeError{
    Time: time.Now(),
    Text: text,
    }
}

func testFunc(i int) error {
// although NewTimeError returns a TimeError type,
// testFunc's return type is error
    if i == 0 {
        return NewTimeError(`parameter in testFunc is 0`)
    }
return nil
}

func main() {
    if err := testFunc(0); err != nil {
        fmt.Println(err)
    }
}
```

Получим примерно следующее:

2021/05/28 11:33:00: Параметр в testFunc равен 0

Если ошибкой может выступать переменная любого интерфейсного типа error, значит, можно использовать операцию утверждения типа (type assertion) для конвертации ошибки в конкретный базовый тип.

if err := testFunc(0); err != nil {
    if v, ok := err.(TimeError); ok {
        fmt.Println(v.Time, v.Text)
    } else {
        fmt.Println(err)
    }
}

Если ошибки могут быть разных типов, логично использовать конструкцию выбора типа:

if err := testFunc(0); err != nil {
    switch v := err.(type) {
    case TimeError:
        fmt.Println(v.Time, v.Text)
    case *os.PathError:
        fmt.Println(v.Err)
    default:
        fmt.Println(err)
    }
}

Но лучше применить функцию As пакета errors, так как она, в отличие от type assertion, работает с «обёрнутыми» ошибками, которые разберём ниже. As находит первую в цепочке ошибку err, устанавливает тип, равным этому значению ошибки, и возвращает true.

if err := testFunc(0); err != nil {
    var te TimeError
    if ok := errors.As(err, &te); ok { //  Сравниваем полученную и контрольную ошибки. Сравнение идёт по типу ошибки.
        fmt.Println(te.Time, te.Text)
    } else {
        fmt.Println(err)
    }
}

Возвращение ошибки не всегда означает, что ситуация критическая. Ошибка может сообщать о статусе или состоянии какого-то действия или ресурса. Например, при проверке наличия файла нужно дополнительно проверить полученную ошибку функцией os.IsNotExist. Другой пример — чтение из источника должно продолжаться до получения ошибки io.EOF, которая сигнализирует о том, что все данные прочитаны.

if _, err := os.Stat(filename); err == nil {
    // файл существует
} else if os.IsNotExist(err) {
    // файл не существует
} else {
    // в этом случае непонятно, что случилось, и нужно смотреть текст ошибки
}

func main() {
    if data, err := ReadTextFile(`myconfig.yaml`); err != nil {
        if os.IsNotExist(errors.Unwrap(err)) {
            fmt.Println(`Файл не существует!`)
        }
    } else {
        fmt.Println(data)
    }
}

В данном примере можно использовать функцию Is(err, target error) bool из пакета errors, которая определяет, содержит ли цепочка ошибок конкретную ошибку.

func main() {
    data, err := ReadTextFile("myconfig.yaml")
    if errors.Is(err, os.ErrNotExist) {
        fmt.Println("Файл не найден")
        return
    }
    fmt.Println(data)
}
