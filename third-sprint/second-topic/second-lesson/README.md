# Second lesson

## Interfaces in the standard library

In the previous lesson, you learned about interfaces in Go. These included `Stringer, Reader, and Writer`, which are already in the standard library, so importing them into production code eliminates the need to write them from scratch. We'll cover other popular interfaces from the standard library below.

### fmt.Stringer

```go
type Stringer interface {
    String() string
}
```

This interface is often used when you need to log a complex object in a single line. The interface definition is in the [fmt](https://golang.org/pkg/fmt/#Stringer) package.

For example, let's take the `User` struct and add an implementation of the `fmt.Stringer` interface to it:

```go
type User struct {
    Email        string
    PasswordHash string
    LastAccess   time.Time
}

func (u User) String() string {
    return "user with email " + u.Email
}

func main() {
    u := User{Email: "example@yandex.ru"}
    fmt.Printf("Hello, %s", u)
}
```

output:

```
Hello, user with email example@yandex.ru
```

The `fmt.Printf` used the implementation of the interface

### Package io

The io package is designed to implement I/O facilities, but it has several convenient interfaces that are used for other purposes as well.

#### io.Reader

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

This interface describes reading from any data stream: network, file system, or buffer. The interface definition is in the [io](https://golang.org/pkg/io/#Reader) package.

The `Read` method reads data from the source into the passed slice of bytes. The source can be any data described by the type. This means we read their structures and write them to bytes. The number of bytes read is implicitly determined by the buffer size—the slice length.

Let's explain the interface's capabilities with an example. Given a buffer, we need to read bytes from it. The [strings](https://pkg.go.dev/strings#Reader) package contains the `strings.NewReader` function, which wraps a regular `string` in a `strings.Reader` structure. This structure has a `Read` method, meaning it implements the `io.Reader` interface:

```go
s := `Hodor. Hodor hodor, hodor. Hodor hodor hodor hodor hodor. Hodor. Hodor!
Hodor hodor, hodor; hodor hodor hodor. Hodor. Hodor hodor; hodor hodor - hodor,
hodor, hodor hodor. Hodor, hodor. Hodor. Hodor, hodor hodor hodor; hodor hodor;
hodor hodor hodor! Hodor hodor HODOR! Hodor hodor... Hodor hodor hodor...`

// Let's wrap the string in strings.Reader
r := strings.NewReader(s)

// Let's create buffer for 16 bytes
b := make([]byte, 16)

for {
    // / strings.Reader copies 16 bytes to b
    // the last pointer is stored in the structure,
    // meaning the next call will copy the next portion of 16 bytes
    // the method also returns the number of bytes read (n) and the error (err)
    // when we reach the end of the string, the method returns the error io.EOF
    n, err := r.Read(b)

    // When working with the io.Reader interface, you should first check
    // n > 0, then err != nil
    // There may be situations where part of the data was read
    // and saved to the buffer, and then an error occurred
    // In this case, both n > 0 and err != nil will be true.
    if n > 0 {
        // выведем на экран содержимое буфера
        fmt.Printf("%v\n", b)
    }

    if err != nil {
        // if it is the end, exit from loop
        if errors.Is(err, io.EOF) {
            break
        }

        // handle reading error
        fmt.Printf("error: %v\n", err)
    }
}
```

The convenience of `io.Reader` is that its user doesn't even need to know where the data comes from: from a file, the network, or generated on the fly. The interface describes a unified method for working with it.

To demonstrate, let's implement a random data generator which will be given in `exmaple-randbyte/`.

### io.Writer

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

This interface allows writing to any possible data stream: a network socket, a file, or a buffer. The interface definition is in the [io](https://golang.org/pkg/io/#Writer) package.

This interface is the opposite of `io.Reader`. It allows writing the byte slice passed to it somewhere. The exact location is implementation-defined.

For example, let's build a large string from substrings, but not using the `+=` operator, because then each iteration would create an extra copy of the entire string. The [strings](golang.org/pkg/strings/#Reader) package includes the `strings.Builder` structure for building a string without unnecessary copying. This structure has a `Write` method, meaning it implements the `io.Writer` interface:

```go
// create strings.Builder
w := strings.Builder{}

for i := 0; i < 50; i++ {
    // fmt.Fprintf gets as arg io.Writer. This allows you to write formatted output
    fmt.Fprintf(&w, "%v", math.NaN())
}

w.Write([]byte("... BATMAN!"))

// print the string
fmt.Printf("%s\n", &w)
```

Let's give an example implementation of the `Write` interface. Suppose we want to calculate a hash from some byte array or sets of arrays. For simplicity, we'll use a simplified hashing function:

```go
package hashbyte

import "io"

type Hasher interface {
    io.Writer // we emedded io.Writer in our interface, to set a requirement for the presence of the Write method
    Hash() byte
}

type hash struct {
    result byte
}

func New(_init byte) Hasher {
    return &hash{
        result: _init,
    }
}

// Write — an array of bytes of any length can be written here, for which the hash will be calculated.
func (h *hash) Write(bytes []byte) (n int, err error) {
    // refresh hash for each byte written to hasher
    for _, b := range bytes {
        h.result = (h.result^b)<<1 + b%2
    }
    return len(bytes), nil
}

func (h hash) Hash() byte {
    return h.result
}
```

usage:

```go
func main() {

    generator := randbyte.New(time.Now().UnixNano())
    buf := make([]byte, 16)

    for i := 0; i < 5; i++ {
        n, _ := generator.Read(buf)
        fmt.Printf("Generate bytes: %v size(%d)\n", buf, n)
    }

    hasher := hashbyte.New(0)
    hasher.Write(buf)
    fmt.Printf("Hash: %v \n", hasher.Hash())

}
```

### Utility functions for io.Reader and io.Writer

#### io.Copy

```go
func Copy(dst Writer, src Reader) (written int64, err error)
```

The function copies all bytes from an `io.Reader` to an `io.Writer`.

Data will be read until the `Read` function returns an error as its second argument. If the error value is `io.EOF`, the function will terminate without error. The number of bytes will also be returned.

**io.EOF** stands for "end of frame"—historically, this was the name of a special character that signified the end of a file.
Let's give a simple example. Let's write a function that copies the contents of one file to another:

```go

func CopyFile(srcFileName, dstFileName string) error {
    srcFile, err := os.Open(srcFileName)
    if err != nil {
        return err
    }
    dstFile, err := os.Create(dstFileName)
    if err != nil {
        return err
    }
    n, err := io.Copy(dstFile, srcFile)
    if err != nil {
        return err
    }
    fmt.Printf("Copied %d bytes from %s to %s", n, srcFileName, dstFileName)
    return nil
}
```

The `os.File` structure implements the `io.Reader and io.Writer` interfaces.
It would be simple to read the entire source file into memory and then copy it to a new one. But what if the source file is hundreds of gigabytes? `io.Copy` works smarter, reading and writing data in small chunks, so it's recommended for such operations.

#### io.CopyN

```go
func CopyN(dst Writer, src Reader, n int64) (written int64, err error)
```

This function copies all bytes from an `io.Reader to an io.Writer`, but no more than `n` bytes. It's the same as `Copy`, but with a limitation: it can be used with data sources that are very large or even infinite. For example, let's write a function that will save data from our random number generator to a file.

```go
// Dump — saves the calculated data to a file.
func (g generator) Dump(n int64, dst *os.File) error {
    _, err := io.CopyN(g, dst, n)
    return err
}
```

If we used `Copy`, the program would continue to run until the disk is full.

#### io.ReadAll

```go
func ReadAll(r Reader) ([]byte, error)
```

The function reads all bytes from the `io.Reader`. Reading ends when the `io.Reader` returns `io.EOF`.

#### io.ReadAtLeast

```go
func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)
```

The function reads bytes from `io.Reader` into `buf` with a limit: if fewer than `min` bytes are read, the error `io.ErrUnexpectedEOF` is returned. This is used when parsing binary data to ensure that the required minimum number of bytes is read.

#### Other io package interfaces

We've provided examples of the basic methods for working with I/O functions and interfaces. There's much more to discover in the package. We recommend opening the io package [documentation](https://golang.org/pkg/io/) to see the definitions of the remaining interfaces.
