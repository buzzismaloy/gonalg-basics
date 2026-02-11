# Third lesson

## Modules

According to the *official documentation*, a **module** is a collection of packages with a common versioning and release cycle. Modules can be loaded either directly from version control systems or from module proxies.
Simply put, it's a group of packages that are stored and updated together. Even your application is a module.
Module metainformation is contained in the `go.mod` file in the module's root directory. A full list of all the directives in this file can also be found in the [documentation](https://go.dev/ref/mod#modules-overview).
Module metainformation describes how it will be built, exported, and which external dependencies it will use.
In this lesson, we'll cover several basic directives using the example of creating local modules, meaning those located in our file system. We'll cover other types of modules in the next lesson. This example is based on Go 1.17, so some aspects may differ in other versions.

### Creating a module

Create a new directory called ypmodule:

```
mkdir ypmodule
cd ypmodule
```

Initialize the module inside the directory using the standard `go mod` utility:

```
go mod init ypmodule
```

A file called `go.mod` is created in the directory, containing:

```
module ypmodule

go <version>
```

The `module ypmodule` line contains the module's import path—this is the prefix relative to which all packages in this module will be imported. For example, to import the package `somepackage` from the `ypmodule` module, you would add the following line to the code:

```go
import "ypmodule/somepackage"
```

It should be noted that in most cases, the `go.mod` file is not edited manually, but rather modified using `go mod`.
In the example, the import path is the single word `ypmodule`. However, the path could also be a full URL, such as `github.com/someuser/somerepo`. Such paths will be discussed later.
The line `go <version>` indicates the Go version used to create this module.

Let's create a `calc` package in the module for working with numbers and place the `math.go` file, which contains the integer addition function, within it.

```go
package calc

func AddInts(a, b int) int {
    return a + b
}
```

In the example, a separate directory has been created for the `calc` package. But if you don't want to add more than one package to the module, you can write the code directly in the directory with the `go.mod` file.

To test the module's functionality, let's create another module next to the `ypmodule` — `main`. The file structure will look like this:

```
.
├── main
└── ypmodule
    ├── calc
    │   └── math.go
    └── go.mod
```

To create a new module, run the command in `main`:

```
go mod init main
```

Create the `main.go` file in the new module, containing:

```go
package main

import (
    "fmt"

    "ypmodule/calc"
)

func main() {
    fmt.Println(calc.AddInts(1, 2))
}
```

Let's try to run the `main` function. And we get an error:

```
main.go:6:2: package ypmodule/calc is not in GOROOT (/usr/local/go/src/ypmodule/calc)
```

***The fact is that the `main/go.mod` file does not describe where to look for the `ypmodule` module. At first, Go went to `GOROOT` and did not find it. Then Go saw that `ypmodule` doesn't look like a URL, so it makes no sense to search for this package on the web.***

Since we are currently working with a local module (that is, its code lies only on our file system), we need to use the ([replace](https://golang.org/ref/mod#go-mod-file-replace) directive to determine its position on the local disk. After adding it, the `main/go.mod` file will look like this:

```
module main

go <version>

// using the replace directive, we specify the position of the root
// of the ypmodule module relative to main/go.mod
replace ypmodule => ../ypmodule
```

***Since `ypmodule` contains other packages and dependencies inside it, they also need to be specified. You may not know about the dependencies and structure of this module, but the Go toolkit will come to the rescue.***

Run the `go get ypmodule` command:

```
go get ypmodule

go get: added ypmodule v0.0.0-00010101000000-000000000000
```

This line appeared in `go.mod`:

```
require ypmodule v0.0.0-00010101000000-000000000000 // indirect
```

It specifies which specific version of the `ypmodule` the `main` module will use during the build. The `// indirect` comment suggests that the `ypmodule` package itself is not imported in the code, only `calc`.

### When to apply the modules?

In practice, always. Modules give the developer a lot of opportunities to manage program dependencies.
