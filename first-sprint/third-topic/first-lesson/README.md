# First lesson

## if-else

```go
if a == 1 {

} else if a == 2 {

} else {

}
```

Here are examples of operators with different conditions:

```go
// logical not
a := false
if !a {}

// logical and
var a, b int
if a == 1 && b == 2 {}

// XOR
var a, b bool
if (a || b) && !(a && b) {}
```
**General rules for conditional operators**:

    - It is mandatory to use curly braces `{ }` to indicate the scope of the operator.
    - It is not necessary to enclose the main condition in parentheses `( )`, but it is more convenient to read the code with them.
    - You can add parentheses `( )` to group the parts of the condition.

Go uses a "lazy" condition check: it goes from left to right until the first `false` and stops because there is no point in checking further. The example of lazy condition check is given in `lazy-check.go`

The `if` statement can consist of two components: initialization and the main condition. This technique allows you to declare a local variable that is used only within the scope of `if`. This can be useful, for example, when you need to convert data for comparison. The example is given in `example.go`

## switch-case

```go
var a int

switch a {
case 1:
    fmt.Println("1")
case 2:
    fmt.Println("2")
case 3, 4:
    fmt.Println("3 or 4")
default:
    fmt.Println("Default case")
}
```
The main switch condition may not be set explicitly:

```go
var a int

switch {
case a == 100:
    fmt.Println("EQ 100")
case a > 0:
    fmt.Println("GT 0 AND NEQ 100")
case a < 0:
    fmt.Println("LT 0 AND NEQ 100")
}
```

In this example, the order of the `case` conditions is important: if you move `a == 100` to the end, then `a > 0` will always trigger first for positive numbers. This form of the operator is similar to the plural `else if`.

Inside the `switch`, you can declare a local variable that is accessible only within the scope of the operator:

```go
a := 6
switch b := a % 5; {
case b == 0:
    fmt.Println("Кратно 5")
default:
    fmt.Printf("Остаток от деления на 5: %d", b)
}
```

To terminate the execution of a `case` prematurely, the `break` keyword is used. This can be useful when there are conditional constructions inside the `case`. In Go, there is no need to explicitly specify a `break` at the end of each `case`, as the next `case` block will fail automatically if the condition is met.

When you need to complete the next block, use the keyword `fallthrough`. If you specify it at the end of the code block, then the block in the next `case` or `default` will be executed after it. The example is given in `example1.go`

The keyword `fallthrough` has features:

    - it can only be used in the last line of the `case`, otherwise there will be a compilation error.;
    - it ignores the condition of the next `case` in order.
