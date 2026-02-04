# Second lesson

Example of declaring custom type:

```go
type Name string
type Fruit string

var fruit Fruit
var name Name

fruit = "Apple"
name = fruit // typing error
```

To fix this error you need to explicitly bring `fruit` to the `Name`.

You can define methods for custom types (as for classes in OOP, more on this later in the course).

You cannot define methods for built-in types in Go.

```go
// declaration of custom type
type MyType string

// declaration of method for the custom type
func (mt MyType) MethodForMyType() {
    // logic of the method
}
```

## Type conversion

To convert one type to another, Go uses the following syntax: `type(variable)`. Let's illustrate with the previous example:

```go
type Name string
type Fruit string

var fruit Fruit
var name Name

fruit = "Apple"
name = Name(fruit)
```

## Aliases

Go also has aliases — don't confuse them with definitions. Aliases allow you to refer to a type in the code by a different name. So, the `rune` and `byte` types represent aliases to `int32` and `uint8`:

```go
type rune = int32
type byte = uint8
```

When defining an alias, there is an assignment mark after the type name:

```go
type MyString = string // MyString is alias to string

var a string
var b MyString

a = b // there is no error
```

You can mix aliases and original types in a single expression.

Aliases were introduced into the language at Google's insistence to solve the problems of a large company that owns large, contiguous, overlapping code bases and arrays of code. The use of aliases facilitates large-scale refactoring of a large volume of already written code from various sources. When writing "fresh" code from scratch, it is better to do without aliases.

## Default values

Unlike the C language, of which Go is the ideological heir, any declaration of a variable and any allocation of memory are accompanied by initialization of this memory. If in C the declared variable can contain a random value remaining in the memory allocated to it, then in Go the variable is first guaranteed to receive a zero value for its type.


All types have default values that automatically initiate the declared variable, unless this has been done explicitly.:

- for `bool`, the default value is `false`;
- for numeric types — 0;
- for reference types — `nil` or an empty pointer.;
- for `string`, an empty string of length 0.

```go
var str string
// both ways to define if string is empty are valid
if str == "" {
   // ...
}
if len(str) == 0 {
   // ...
} 
```
