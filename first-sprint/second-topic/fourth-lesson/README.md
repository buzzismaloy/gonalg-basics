# Fourth lesson

## Scope

In Go, the scope can be characterized by two parameters:

    - globality/locality;
    - exportability/non-exportability.

### Globality/locality

Variables, constants, and functions declared in the function body are characterized as local, meaning their scope is limited to the block of code in which they are declared.
In Go, local variables can be declared with the keyword var or the short notation :=.
Let's look at an example that illustrates the scope of a local variable which is given in `example.go`

The entity functions declared outside the body are characterized as global and are in the scope of all code blocks in the package. In other words, global objects are available in all files located in the same directory.
Global variables cannot be declared using short notation :=, the var keyword is needed. If a global variable is not initialized, it takes a null value of the specified type.

The example is given in `example1.go`

## Exportability/non-exportability

If the previous property characterizes the accessibility of an object within a package, then exportability regulates accessibility from other packages. A package in Go is all files with an extension `.go` in the same directory, except for those ending in `_test.go`.

**Only global entities can be exported.**

**Exported** variables, constants, and functions of a particular package are available from those packages that import this package with the `import` keyword.

The exported objects are accessed by the `<package name>.<Name of the entity>` construct. Let's illustrate with the example of `Hello, world!`:

```go
package main

import "fmt"  // importing the fmt package of the standard library

func main() {
    fmt.Println("Hello, world!")
}
```

Here, in the `main` function, the program accesses the `fmt` standard library package by calling its exported `Println` function.

**How does Go determine that an entity is exported? If the name of a variable, constant, or function
begins with a capital letter, then it is exported. If it is lowercase, then it is not exported.**

**Non-exported** entities cannot be accessed outside the package. This approach is part of the special OOP that implements data hiding or `encapsulation` in Go.

Let's use an example to illustrate the definition of exported and non-exported objects which is given in `example-project` directory.
