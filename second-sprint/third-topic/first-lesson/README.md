# First lesson

## Packages and modules

### Package

A **package** is a unit of compilation, namespace, and import. All the Go code is in some kind of package.
Simply put, a package is a set of source code files that are located in one project folder. Packages allow you to logically divide your project into components. All code elements (types, constants, variables, functions) are available inside the package as if they were declared in a single file. In each file, the Go source code must begin with the package declaration: the `package` keyword and the package name.

```go
package main

import "fmt"

func main() {
    fmt.Println(`Hello, world!`)
}
```

All files in the directory must have the same package name, because the compiler will process all files in the directory at once (unless you specify a specific one) and an error will occur with different package names. However, there is an exception to this rule — files whose name ends with `_test.go`, may have a different package name, as they are ignored when building the program.
The "one directory, one package" approach simplifies the work and further program support. It is strongly recommended that the package name match the folder name, as they will be easier to work with.
If you need to create an executable file, the package must have the name main and the function of the same name in one of the files. In all other cases, if the package is to be used as a library, it is desirable that the package name reflect its purpose.

### The best name for a package and its elements

1. Write the package name in lowercase letters. It's good if it's short but informative: other programmers should understand what this package does.
2. Choose a unique package name within the repository. Otherwise, when importing multiple packages with the same name, you will have to use aliases.
3. Do not write common words in the package name: util, base, tools, lib, common. Usually, these names don't tell you anything about the real purpose of the package and only litter the namespaces. Remember that the package name will be used in the code.
4. Do not use the plural. But there are exceptions: the strings, bytes, errors packages in the Go standard library are named in such a way as to avoid conflicts with types. Sometimes it is acceptable to use the plural, for example, if you need to show that the handlers package contains several handlers.

***When naming package elements, remember the main rule: when importing a package, only those definitions whose names begin with a capital letter are available. In the example below, when importing a package, you can call mypkg.Process(), but not mypkg.calculate(). This rule applies to functions, methods, constants, types, and global variables.***

```go
package mypkg

func Process() {
    ...
}

func calculate() {
    ...
}
```

Do not use words from the package name when naming exported objects. For example, if there is an `md` package for working with the Markdown format, the HTML conversion function can be called `HTML`. You should not name the `MarkdownToHTML` function: the package name already indicates what you are working with.

A **module** is a complete library or application that can contain internal packages and import external ones. You can place the module in any directory, regardless of the Go installation path. If your package is a separate library or application, design it as a module and place it wherever it suits you.
There are no strict rules for organizing directories, but it is better to start the path with the domain or the author's name in order to avoid name conflicts when using packages publicly. For example, you can create a `golang` directory and place all Go projects there, create a separate subdirectory for private closed projects, and place public ones immediately in the subdirectory of the corresponding repository. In this case, the packages from GitHub will be located in the following subdirectories `github.com/author/packagename`.

## Export

As mentioned above, all code elements are available inside the package. But how to make them accessible from the outside? There is an export concept in Go for this.

***Any element (type, constant, variable, function) is exportable, that is, available to external packages for import if its name begins with a capital letter. Otherwise, such an element is only available inside the package. This is akin to the `public` and `private` modifiers from C++/Java or `static` from C++***

Examples of exportable elements:

```go
var ParsedString string

func Print (s string) {}

const Red = 3

type MyStruct struct {
    a int
}
```

Examples of non-exportable elements:

```go
func someTestFunc()

const green = 5
```

The exported elements represent the external interface of your package. Let it be minimalistic. If an item is not needed outside the package, it is better to make it non-exportable.
You should also be very careful about exporting variables: it may be better to make functions that change your variables so that the code does not become inflexible and dangerous. It will be more difficult for you to make changes to the package code if the external code uses variables from your package.

## Import

In order for one package to use another, it must be imported. Importing a package is somewhat similar to a similar process in Python. It is performed using the `import` keyword.
You can specify `import` for each package, but more often they list all packages inside parentheses. The `import` should go immediately after the `package` is declared.

```go
package main

import "fmt"
import (
    "encoding/json"
    "strings"

    "github.com/yuin/goldmark"
    "golang.org/x/crypto/bcrypt"
    "gopkg.in/yaml.v2"
)
```

***Packages can be divided into groups using empty lines: for example, the standard library, internal packages, and external packages. In this case, it is better to sort the packages within the group alphabetically. If a package has several versions, you can specify the required version after the package name.***

As mentioned above, only exported variables are imported. If the packages form a loop during import, that is, they import each other directly or through other packages, the compiler will give an error. For example, if we have three packages, then such an import sequence is unacceptable — the compiler will simply fall into a loop, bypassing the packages.:

```go
package A

import "B"

// ...

package B

import "C"

// ...

package C

import "A"
```

When accessing imported objects, the package name and the element name must be separated by a dot: `fmt.Println(...), yaml.Marshal(...)`.

### Renaming the import

There may be a situation when you need to import two packages with the same name: for example, `crypto/rand` and `math/rand`. Then you need to specify a different alias for one of the packages. In addition, an alternative name during import can be useful if the package has a long name. As for the new package name, it is indicated by a space before the full name.

```go
import (
    hl "github.com/yuin/goldmark-highlighting"
)
```

Agree, the challenge is `hl.NewHighlighting(...)` it does not clutter up the code, unlike `goldmark-highlighting.NewHighlighting(...)`. However, you should not get carried away with renaming without obvious need.

### The order of import

Let's take a closer look at how the packages are imported.

    1. When compiling a program, the compiler starts with the `main` package. If there are imports of any packages in `main`, it goes to them and compiles them until all the necessary packages are compiled to build the program.
    2. Then the compiler compiles the `main` package and builds the main application, and then the following happens during the execution of the program:
        1. The package variables will be initialized in the order in which the packages were imported.
        2. After that, the `init()` functions inside each package will be executed. There may be several `init()` functions, and they will be executed in the order in which they were declared.
        3. And after their execution, it will be the turn of the `main` function.

Thus, part of the initialization code can be executed even before the main program is started.

### Empty import

If you import a package but do not use it inside a file, the compiler will return an error. However, there are situations when an imported package needs to call the `init` function to initialize data. In this case, the underscore `_` is used instead of the alternative name of the imported package.

As an example, there is a very convenient `embed` package from the standard library. It allows you to initiate the values of string variables with the contents of a file.

```go
import _ "embed"

//go:embed insert.sql
var queryInsert string
```

In the first case, we can get an empty import when the package is not used directly, functions from it are not called, and only annotations in comments are imported.

The second case is an application in the development process. Go prohibits compiling code that contains unused imports, but sometimes you need to leave imports when debugging. Then it can be marked as empty. Naturally, this code is used only for debugging purposes and it's worth keeping it away from your codebase.

## Code organization

Go does not impose restrictions on the project structure and package layout. You can organize files in a project any way you want, but there are recommendations that you can follow to make your project understandable to other programmers.
There is an approach called **Standard Go Project Layout**. In it, the project code is organized in the form of the following directories:

**cmd**

If there are several binary files in the project, create subdirectories for them in `cmd`. The names of the subdirectories must match the names of the executable files.

```
- cmd
    - client
        main.go
    - server
        main.go
```

If there is only one executable file in the project, the source files can be placed directly in the project directory. Try to use one or two source files. The rest of the code needs to be placed in separate packages.

**internal**

The `internal` directory contains the internal packages of the Go project. At the compiler level, it is prohibited to import such packages from outside the parent directory `internal`. For example, the package `.../root/client/internal/a/b` can only be imported in a directory tree file starting with `.../root/client`, and cannot be imported into `.../root/server` or another repository.

It is recommended to place the entire main source code of the program in the `internal` directory, divided into subdirectories with the corresponding packages. Depending on the complexity of the project, packages may have different levels of nesting.

**pkg**

Define the `pkg` directory for packages that can be used in other projects. It is preferable to create separate repositories for public projects.

**vendor**

The `vendor` directory contains external packages. With the advent of Go modules, all dependencies are stored in the module cache. Therefore, the `vendor` directory can be used on older versions of Go or if you want to be sure that all dependencies are located inside the project directory.

**test**

As a rule, each package contains one or more `name_test.go` test files. The `test` directory can be used for comprehensive testing with additional tools.

**docs**

The `docs` directory is used for project documentation. This can be documentation for users or an addition to the documentation that `godoc` automatically generates.

**Other directories**

Here are more directory options that are found in Go projects.:

    - `api` — additional files for services with API.
    - `assets` — additional resource files. For example, pictures.
    - `build` files for packaging and continuous integration.
    - `configs` are configuration files.
    - `deployments/deploy` are configuration files and templates for services, operating systems, and containers.
    - `examples` — examples of using applications and libraries.
    - `scripts` — scripts for installation, configuration, and other actions with the project.
    - `tools` — tools for project support. They can be written in Go using project packages.
    - `website` is a directory with files for the project's website.

This is an incomplete list. You can create directories with other names. Try to give names that would reveal the purpose of the directory.
