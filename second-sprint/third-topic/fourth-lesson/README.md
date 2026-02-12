# Fourth lesson

## External dependencies

**External dependencies** are packages that are not part of the standard library.

In addition to packages stored in the local file system, you need to work with packages from external sources. For example, you may need a framework or a specific package.
There are two ways to download packages in Go:

- manually install the `go get` [utility](https://pkg.go.dev/cmd/go/internal/get);
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

From the previous lesson, you learned that the module system allows you to explicitly list a project's dependencies. This is done using the [require](https://go.dev/ref/mod#go-mod-file-require) directive. After running the code, Go will automatically download the packages listed in the `require` block and cache them in the `$GOPATH/pkg/mod directory`.

For example, if a module has a single dependency on the library `github.com/stretchr/testify`, the `go.mod` will look like this:

```
module somemodule

go <version>

require github.com/stretchr/testify v1.7.0
```

A special file, `go.sum`, will also be created to ensure 100% reproducibility of runs. It contains the hash sums of all modules, thus guaranteeing reproducible installation of modules in different environments. You can read more [here](https://github.com/golang/go/wiki/Modules#is-gosum-a-lock-file-why-does-gosum-include-information-for-module-versions-i-am-no-longer-using).

***Manually specifying dependencies in `go.mod` is often unnecessary. Go can automatically update the dependency list in `go.mod` when running a program (e.g., calling `go run, go build, or go test`) if the library import path is the URL to the code repository and no other version of the library is required.***

But what if you don't need the latest version of the library? In this case, you can specify the required version in the `require` directive. In the example, it was `v1.7.0`.

## Versioning

Go uses [Semantic Versioning](https://semver.org/spec/v2.0.0.html) (semver).

Semantic versioning is a generally accepted format for numbering releases (packages, modules, libraries, etc.). The version is written in the format `vX.Y.Z`, where:

- `X` is the major version,
- `Y` is the minor version,
- `Z` is the patch version.

For example, a package with version `v1.2.3` has a major version 1, a minor version 2, and a patch version 3.

- The patch version should be changed if minor fixes are made to the package.
- The minor version is changed if new functionality is added.
- The major version is changed if backward compatibility with the previous version of the code is broken.

The incrementing of the high-order bit should be accompanied by the zeroing of the low-order bits. That is, incrementing the major version resets the minor and patch versions, while incrementing the minor version resets only the patch version.
A semantic version can have a pre-release suffix separated by a hyphen. For example, `v1.2.3-beta`.

When working with Git, a package version is determined by a release tag (you can read more about tags [here](https://git-scm.com/book/ru/v2/%D0%9E%D1%81%D0%BD%D0%BE%D0%B2%D1%8B-Git-%D0%A0%D0%B0%D0%B1%D0%BE%D1%82%D0%B0-%D1%81-%D1%82%D0%B5%D0%B3%D0%B0%D0%BC%D0%B8)). Tags are simply metadata. There are two types of tags: lightweight and annotated. The former simply mark a commit with a version, while the latter contain a wealth of additional information, such as the author's name, creation date, checksum, and so on.
When working with Git, the package version is determined by the release tag (you can read more about tags here). Tags are simply metadata. There are two types of tags: lightweight and annotated. The former simply mark a commit with a version, while the latter contain a wealth of additional information, such as the author's name, creation date, checksum, and so on.

If you need to increase the package version to a certain value (for example, `v1.2.3`), you need to run the following commands:

```
git tag v1.2.3
git push --tags
```

## Substituting Dependencies

Sometimes you need to replace a library in your code with a fork (copy) of it, but without changing all import paths. For example, if a critical bug is discovered in the library, a PR has been submitted, but there's no time to wait for it to be merged.
The replace directive comes to the rescue. In the previous lesson, it was used to determine the location of a local module in the file system. But this directive also allows you to replace one external module (or a specific version of it) with another.

This can be done like this:

```
replace (
    golang.org/x/net v1.2.3 => example.com/fork/net v1.4.5
    golang.org/x/net => example.com/fork/net v1.4.5
)
```

## Maintaining Dependencies

So, modules allow you to comprehensively describe which libraries a package depends on, allowing developers to download all required dependencies.
But what if the code depends on a repository that was suddenly deleted or made private?

Having to access the corresponding version control system repositories each time is a [problem](https://qz.com/646467/how-one-programmer-broke-the-internet-by-deleting-a-tiny-piece-of-code). The repository may be unavailable due to locks, or it may be a local repository that needs to be connected to via a proxy, or the maintainer may simply delete their repository for their own reasons.
There are two solutions: vendoring and `go proxy`.

### Vending

Vending is the practice of storing the source code of dependencies directly within the module.

***In Go, vendored libraries should be stored in the `vendor` folder next to the `go.mod` file. To vendor all used libraries, run the `go mod vendor` command.***

But vendoring has its drawbacks:

- The `vendor` folder isn't used by default unless the `-mod=vendor` flag is set when running the program.
- In the case of large monorepos, this folder grows uncontrollably, slowing down repository cloning (which happens constantly during automated tests and other CI/CD runs).
- Pull requests, which consist of 99% vendored files, are difficult to review.

Most of these problems can be solved with module proxies.

### Module Proxy Servers

A module proxy server is a dedicated server for storing downloaded and compiled modules.

Using a module proxy server has several advantages:

- The `vendor` folder is no longer needed. This eliminates the need for unnecessary files in the repository, preventing them from polluting the diff during review.
- Proxies significantly speed up downloads compared to version control clients.

#### Using proxy servers

The module proxy server address is specified by the `GOPROXY` environment variable. This variable's value is a comma-separated list of proxy server addresses or the special value `direct`.

For example:

```
GOPROXY=https://proxy.golang.org/,direct
```

When downloading dependencies, `go get` follows the URL specified by this variable and attempts to download the required dependency. If no server has the required dependency, or the special value `direct` is encountered, `go get` will go directly to the corresponding version control system repository.
The default proxy server is `proxy.golang.org`. This public proxy server is maintained by Google and continuously caches all modules downloaded through it.

## Checksum

The `go.sum` file stores information for 100% reproducible builds. However, there's a problem: at the time of the first build, there's no information about module checksums, and at that point, a malicious proxy server could send us modified code. Checksums are simply hashes that can be used to determine whether a module's contents have changed.

To address this issue, the Go ecosystem has a module checksum database — [checksum](https://go.googlesource.com/proposal/+/master/design/25530-sumdb.md).

When downloading a new module, Go consults the configured checksum database and checks whether the downloaded code's checksum matches the one stored in the database. If it doesn't, Go will generate an error and terminate.
If the `go.sum` file already contains the `checksum` of the module being downloaded, the checksum will not be used.
The address of the `checksum used is determined by the `GOSUMDB` environment variable (by default, `sum.golang.org`).
You can also completely disable `checksum` use by setting `GOSUMDB=off`.

## Private Dependencies

Often, code under development uses both public and private modules. It's important to prevent private modules from leaking into a public dependency proxy. For example, you're developing a module for internal company needs and don't want to share the code with competitors.

There are several ways to avoid this:

- Set up your own private proxy. [one](https://github.com/gomods/athens),
  [two](https://github.com/goproxy/goproxy), [three](https://thumbai.app/).
- Set `GOPROXY` to `direct` to always bypass proxy servers.
- Set the `GOPRIVATE` variable. This variable's value is equal to the import path mask for modules (e.g., `GOPRIVATE=*.internal.company.com`) that don't need to use dependency proxies and instead should access version control systems directly.
