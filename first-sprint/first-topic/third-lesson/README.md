# Third lesson

## Small cheat sheet

| Command           | Description                                                                 |
|-------------------|-----------------------------------------------------------------------------|
| go build          | Compiles the packages named by the import paths and produces an executable |
| go run            | Compiles and runs the specified Go program directly                         |
| go install        | Compiles and installs packages and dependencies to `$GOPATH/bin` or `$GOBIN`|

| go get            | Downloads and installs packages and updates module dependencies             |
| go mod init       | Initializes a new Go module in the current directory                         |
| go mod tidy       | Adds missing and removes unused module dependencies                          |
| go mod vendor     | Copies module dependencies into the `vendor` directory                       |

| go test           | Runs tests and benchmarks for packages                                       |
| go tool cover     | Analyzes and displays code coverage results                                   |
| go tool pprof     | Analyzes profiling data (CPU, memory, goroutines, etc.)                      |

| go fmt            | Formats Go source code according to standard style                            |
| go vet            | Examines Go code for suspicious constructs and common mistakes                |
| go env            | Prints Go environment information
