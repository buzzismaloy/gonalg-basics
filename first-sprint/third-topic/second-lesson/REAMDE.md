# Second lesson

## Loops

Omitting the scope of loops, let's consider their forms in the Go language. In programming, the constructions `for`, `while`, `do — while`, and others are used for loops. The creators of Go adhere to the rule "the simpler the better", so for all types of loops in the language there is only one keyword — `for`.

### Infinite loop

```go
for {

}
```

### For loop

```go
v := 0
for i := 1; i < 10; i++ {
    v++
}
fmt.Println(v)
```

The classical form of the loop consists of three components:

    - `i := 1` — initialization (pre-action): it is executed once when entering the scope of the loop;
    - `i < 10` is the main condition: as long as the condition is `true`, the iterations will continue;
    - `i++` — post action: it is executed at the end of each iteration of the loop.

It is not necessary to fill in each component — you can omit it.

There are two more options for an infinite loop, but in the form of three components:

```go
for ;; {}
for ; true; {}
```

The components of the loop can take on a more complex form:

```go
for a, b := 5, 10; a < 10 && b < 20; a, b = a + 1, b + 2 {
    // do stuff
}
```

## While loop

Only one condition can be specified in the `for` loop. In this case, the necessary initial initialization must occur before the loop, and actions affecting the condition must be performed inside the loop.

```go
package main

import "fmt"

func main() {
    i := 0
    for i < 5 {
        fmt.Println(i)
        i++
    }
}
```

## Range loop

```go
array := [3]int{1, 2, 3}
for arrayIndex, arrayValue := range array {
    fmt.Printf("array[%d]: %d\n", arrayIndex, arrayValue)
}
```

The `range loop` is used for complex types — slice and map.

## break & continue keywords

The current iteration of the loop can be interrupted with keywords:

    - `break` — exit the loop;
    - `continue` — moves to the next iteration of the loop (calling a post action, if specified).

For example, let's calculate the sum of all even numbers from 0 to a given limit:

```go
sum, limit := 0, 100
for i := 0; true; i++ {
    if i % 2 != 0 {

    }

    if sum + i > limit {

    }

    sum += i
}
fmt.Println(sum)
```

The keywords `break` and `continue` refer to the loop closest in scope.

## Labels

Labels allow you to jump to different parts of the code.

You can specify a label for operators:

    - `break`;
    - `continue`;
    - `goto` (an unconditional jump operator that allows you to jump to any place in the code).

The example for `break`:

```go
outerLoopLabel:
    for i := 0; i < 5; i++ {
        for j := 0; j < 5; j++ {
            fmt.Printf("[%d, %d]\n", i, j)
            break outerLoopLabel
        }
    }
    fmt.Println("End")
```

```
[0, 0]
End
```

Here, the `break outerLoopLabel` interrupts the execution of the outer loop.

The example for `continue`:

```go
outerLoopLabel:
    for i := 0; i < 5; i++ {
        for j := 0; j < 5; j++ {
            fmt.Printf("[%d, %d]\n", i, j)
            continue outerLoopLabel
        }
    }
    fmt.Println("End")
```

```
[0, 0]
[1, 0]
[2, 0]
[3, 0]
[4, 0]
End
```

Here, the `continue outerLoopLabel` causes the transition to the next iteration of the outer loop. If you replace the `continue outerLoopLabel` with a `break`, the result will be similar.

The keywords `break` and `continue`, without specifying a label, refer to the current (nearest) scope of the code.

The use of labels in Go, as in many other languages, is a topic of eternal debate. It is believed that labels make the code unobvious, turn it into a so-called ***spaghetti code***.
