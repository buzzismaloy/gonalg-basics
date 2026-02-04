# Third lesson

General way of the variable declaration: `var name type = expression`

There are also three more way to declare a variable:

    - `var name = expression`
    - `var name type` / `var name type = expression`
    - `name := expression`

Many variables can be initialized separated by commas or by calling a function that returns multiple
values:

```go
var now = time.Now()  // now equals current time and has time.Time type
var pi, e = 3.1415, 2.7183
var f, err = os.Open("myfile.txt") // os.Open returns 2 values
```

You may not specify `var` before each variable, but combine variables into blocks `var(...)`. This is useful when you need to identify entities that are similar in meaning.

```go
var height int
var length int
var weight float64
var name   string
var company = "Рога и копыта"

// equivalent

var (
    height, length int
    weight float64
    name   string
    company = "Рога и копыта"
)
```

## Short notation

```go
i := 10
f := 5.1
doublef := 2*f // doublef has float64 type and equals 10.2

// equivalent

var i = 10
var f = 5.1
var doublef = 2*f
```

If you need specific type other than default one you can use type conversion as in the example below:

```go
int64Var := int64(5)
float32Var := float32(101.1)

// equivalent

var int64Var int64 = 5
var float32Var float32 = 101.1
```

Short notation works great with multiple declarations. But in this case, at least one variable in the expression must be new. Otherwise, a compilation error will occur.:

```go
pi, e := 3.1415, 2.7183
// when specifying values, you cannot use :=, since
// both variables are already defined
pi, e = 3.14159, 2.71828

f, err := os.Open("myfile.txt")
```

## Constants

### Named constants

Named constants can be initialized with expressions consisting of constants or literals of the following types:

    - numbers;
    - lines;
    - symbols (rune);
    - boolean values.

```go
const pi = 3.14159
const doublePi = pi * 2
const version = "1.0.0"

// equivalent

const (
   pi = 3.14159
   doublePi = pi * 2
   version = "1.0.0"
)

func main() {
    fmt.Println(version, pi, doublePi)
}
```

### Untyped constants

Named constants can be of different types. The type is related to the stored value:

```go
const intConst = 5
const floatConst = 5.0
const runeConst = 'A'
const strConst = "Hello, world!"
const boolConst = true
```

It may seem that if you omit the type when declaring a constant, the compiler will select it itself, as is the case with the short form of variable declarations. This is only partially true. In the case of constants, the lack of explicit type indication is more important.

For example, if you declare an `intConst` constant and assign it the value ***5***, you get an **integer constant with an undefined type (`untyped int`)**. The specific type of value of this constant has not yet been determined and will be interpreted differently by the compiler in different contexts. This allows you to weaken the typing for constants without giving up strong typing globally.

Thanks to this approach, the example given in file `example.go`

If you define `id` as a variable var `id = 100`, compilation errors will occur when defining variables `i` and `f`:

```
./prog.go:10:16: cannot use id (variable of type int) as type int64 in variable declaration
./prog.go:11:18: cannot use id (variable of type int) as type float64 in variable declaration
```

If constants, like variables in Go, always had a specific type, then it would be more difficult to work with them. Moreover, Go allows you to mix numeric literals of different types (untyped int, untyped float), so the following expression (which is also given in `example1.go` is correct:

```go
var a float64
a = 5 + 5.0
```

Constants, like variables, can be grouped:

```go
const Program = "Моя программа"
const Version = "1.0.0"

// equivalent

const (
   Program = "Моя программа"
   Version = "1.0.0"
)
```

If a constant has no value specified in the group, then it is equal to the value of the previous
constant:

```go
const (
    pi = 3.1415
    e
    name = "John Doe"
    fullName
)

func main() {
    fmt.Println("pi =", pi, "e =", e)
    fmt.Println("name =", name, "fullName =", fullName)
}
```

The result of the program will be:

```
pi = 3.1415 e = 3.1415
name = John Doe fullName = John Doe
```

### Typed constants

If you specify the type of a constant explicitly when declaring it, it becomes **typed** and obeys Go's strong typing rules. In this case, you are working with a constant as an immutable variable:

```go
const flag uint8 = 128

func main() {
    var i int = flag
    fmt.Println(i)
}
```
When compiling this example, the error `cannot use flag (constant 128 of type uint8) as type int in variable declaration` occurs, since the constant `flag` has type `uint8`, and the variable `i` has type `int`.


### The iota keyword

What if an enum needs to be implemented in the code? There is no built-in syntactic construct or special type in Go for this. However, you can simply declare a number of constants and work with them which is presented in `example2.go`.

For convenient declaration and initialization of constant blocks in Go, there is an automatic iota increment. When each const block is declared, the iota value is 0 and increases by 1 for each subsequent element which is presented in `example3.go`.

This construction is used not only for enumerations. The iota keyword can also be used in arithmetic expressions to quickly declare a series of values with a progression. It should be remembered that iota increases by one for each line where the name of the constant is specified, even if it has been assigned a specific value.

```go
const (
    _ = iota*10  // обратите внимание, что можно пропускать константы
    ten
    twenty
    thirty
)

const (
    hello = "Hello, world!"  // iota равна 0
    one = 1                  // iota равна 1

    black = iota   // iota равна 2
    gray
)

func main() {
    fmt.Println(ten, twenty, thirty)
    fmt.Println(black, gray)
}
```

The program outputs:

```
10 20 30
2 3
```

You can check this example simply by running `example4.go`.

### Custom types in constants

Suppose you need to define constants for the days of the week.

```go
const (
    Monday = iota + 1
    Tuesday
    //...
    Sunday
)
```

If you list them like this, then all constants will have an untyped numeric type and can be used in any expression, which can be confusing: `var i int = Monday + 1`. In such cases, it is worth defining a custom type and specifying it when defining constants.

The example is given in `example5.go`

### Literals

In Go, you can use different representations of string and numeric literals. Let's illustrate with the example of the integer `1000`:

```
1000
1000.0
1_000 // you can separate the parts of the number with the character '_' for ease of perception
01750 // octal representation, starting from 0
0x3e8 // hexadecimal representation
0b001111101000 // binary representation
```

Any of these literals can be used in expressions and will give the same value.
