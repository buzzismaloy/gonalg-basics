# Second lesson

## Package and file system communication

If you do not use modules (they will be discussed in the next lesson), then the path to import the package is the path to the package directory relative to `${GOPATH}/src`.

When installing the compiler, it creates the `GOPATH` environment variable, which specifies the path to the installation folder.

Let's say we need to create a package for a financial project. Let's call it `finance`.

As we write the code for this package, we discover that the same mathematical functions are needed in different places. So we move them into a separate subpackage, `finmath`. The package tree in the file system will then look something like this (file names may vary):

```
finance
├── money_calculation.go
└── finmath
└── arithmetics.go
```

From the `GOPATH`, it looks like this:

```
GOPATH
└── src
    └── finance
        ├── money_calculation.go
        └── finmath
            └── arithmetics.go
```

Now you can import the `finance` and `finmath` packages using the paths "`finance`" and "`finance/finmath`" respectively:

```go
package main

import (
    "fmt"

    "finance"
    "finance/finmath"
)

func main() {
    fmt.Println(finance.GetMostRecentBill())
    fmt.Println(finmath.Add(1, 2))
}
```

The path to packages hosted in version control systems is usually `${GOPATH}/src/<VCS_URL>/<USER_NAME>/<REPO_NAME>`, and for example, for the [testify](https://github.com/stretchr/testify) testing library, it is `${GOPATH}/src/github.com/stretchr/testify`.

***In Go, cyclic imports are prohibited at the compiler level. If several packages are stuck in a loop,
you should either move the problematic code into separate packages or rethink the application's logic. A
**cyclic import** is a sign of an architectural problem in the application. If the application's
architecture is properly designed, cyclic imports should not occur.***
