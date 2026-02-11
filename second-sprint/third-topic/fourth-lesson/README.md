# Fourth lesson

## External dependencies

**External dependencies** are packages that are not part of the standard library.

In addition to packages stored in the local file system, you need to work with packages from external sources. For example, you may need a framework or a specific package.
There are two ways to download packages in Go:

    - manually install the `go get` utility (detailed documentation can be found [here](https://golang.org/pkg/cmd/go/internal/get/));
    - use the dependency list in `go.mod`.

We will consider both options, but in the future we will use only `go.mod`, since it is better to choose the second method in projects. This will allow you to store dependency descriptions locally and make it easier for other developers to install.

### Installing packages manually using the `go get` utility

If you use `go get`, installing a third-party package looks like this:

```
go get github.com/username/packagename
```

The `go get` utility descends to `https://github.com/username/packagename` and it will download the required package if it was found at the transmitted URL. If the version control system supports multiple protocols, Go will try each one in turn. For example, in the case of git, it will try `https://` and `git+ssh://`.
After that, the downloaded data will be placed in `GOPATH/src/username/packagename`.

### Installing dependencies from `go.mod`

From the previous lesson, you learned that the module system allows you to explicitly list a project's dependencies. This is done using the [require](https://go.dev/ref/mod#go-mod-file-require) directive. After running the code, Go will automatically download the packages listed in the require block and cache them in the $GOPATH/pkg/mod directory.

For example, if a module has a single dependency on the library github.com/stretchr/testify, the go.mod will look like this:

module somemodule

go 1.16

require github.com/stretchr/testify v1.7.0

A special file, go.sum, will also be created to ensure 100% reproducibility of runs. It contains the hash sums of all modules, thus guaranteeing reproducible installation of modules in different environments. You can read more [here](https://github.com/golang/go/wiki/Modules#is-gosum-a-lock-file-why-does-gosum-include-information-for-module-versions-i-am-no-longer-using).
