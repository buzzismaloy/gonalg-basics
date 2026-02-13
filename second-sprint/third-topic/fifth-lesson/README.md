# Fifth lesson

## Rules of good form when supporting your own modules

### Code formatting and gofmt

Code written in a consistent style is easy and pleasant to read. Compare for yourself:

Poorly formatted code:

```go
func Add(a, b ArrInt) ArrInt

{
length := len(a)
if length-len(b) > 0
    {
length = len(b) }
c := make(ArrInt, length)

for i := 0; i <length; i++
    {
    c[i]=a[i]+b[i]
    }

    return c }
```

Well-formatted code:

```go
func Add(a, b ArrInt) ArrInt {
    length := len(a)
    if length-len(b) > 0 {
        length = len(b)
    }
    c := make(ArrInt, length)
    for i := 0; i < length; i++ {
        c[i] = a[i] + b[i]
    }
    return c
}
```

Different languages handle formatting differently. For example, Python has the official PEP 8 standard, while C++ has several popular styles.

Go addresses this issue in an unusual way: all code must be formatted using the standard `gofmt` utility or its extended version, [goimports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports). The `goimports` utility does the same thing as `gofmt`, but also automatically optimizes and sorts imports.

This utility eliminates the need to remember whether to use tabs or spaces when aligning code, how many blank lines to leave between two adjacent functions, and other little details that machines handle much better than humans. Use it before publishing all your packages—other programmers will thank you.
Formatting isn't the only thing that determines code quality. Developers can inadvertently use constructs like` a = append(a,) or fmt.Printf("Here must be number %d")`. In these cases, linters can help.

### Linters and go vet

[vet](https://golang.org/cmd/vet/) is the most widely used linter in the Go ecosystem.

Unlike the `gofmt` code formatter, the linter doesn't provide a 100% guarantee of correctness—false positives and false negatives may occur. You'll likely need to tweak the settings a bit before the linter works correctly in every project. The effort required to configure and integrate the linter into your CI/CD systems will pay off in the time it saves you during code review.

### Documentation and godoc

The Go ecosystem has a standard utility for generating documentation from comments in code: [godoc](https://pkg.go.dev/golang.org/x/tools/cmd/godoc). Run `go install golang.org/x/tools/...@latest` to install all packages and utilities from `golang.org/x/tools`, including `godoc`.

Documentation for any entity (function, structure, variable, or package) is a comment preceding the declaration of that entity. For example:

```go
// Foo выполняет очень важную роль в проекте — ничего не делает :)
func Foo() {}

// описывает новый тип данных «никнейм» на основе стандартного строкового типа
type nickname string
```

Based on these comments, godoc can generate documentation in HTML format, man pages, and more.
Documenting the public API of developed packages allows third-party developers to get answers to questions about functionality without reading the source code. Document your packages, and your users' lives will be better.

### Code testing

As we all know, code without tests is inherently broken. Therefore, code testing is an essential part of development.

### Versioning without surprises

The previous lesson covered the concept of semantic versioning and discussed when to change a package's version: patch, minor, or major. These rules should be followed, especially when breaking backward compatibility. After all, when developing a project, no one wants to suddenly discover that a feature they never knew existed has been removed, but it's causing the compilation of a crucial library to fail.

To protect against outrage from your library's users, remember to increment the major version every time you break backward compatibility in the library.
