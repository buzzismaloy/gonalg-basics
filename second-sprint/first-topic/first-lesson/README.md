# First lesson

## Composite types

### Pointers

The syntax of a pointer type variable is very simple:

```go
var p *int
```

Here we have created a variable of the type "pointer to an integer". In Go, you can create a pointer to any data type.

Physically, a pointer is a memory location that stores the address of the cell that the pointer is "looking at". After creation, the pointer does not "look" at any memory location in the computer and has a zero value. It looks like `nil`.

Для того чтобы присвоить указателю значение (адрес какой-либо переменной), используется операция взятия адреса `&`:

```go
var a int = 5
p := &a

fmt.Println(a,p) //a=5 p=0xc0000b2008
```

To get the value of a pointer, there must be a variable in memory that it "looks at". This value is called **addressable**. It's more difficult with constants — you won't be able to get the address from them.

```go
const c = 5
p1 := &"abc" // error
p2 := &с // error
```

The type of the variable to which the pointer is being created must match the type of the pointer.

```go
var p *int
var a int = 5
var b string = "abc"
p = &a
p = &b // error
```

Composite type literals create a variable of the appropriate type in memory, so you can create a pointer like this:

```go
type A struct {
    IntField int
}

p := &A{
    IntField: 10,
}
```

And Go also has a built-in `new()` function. The type is passed to it as a parameter, and a pointer to a new variable of the corresponding type is returned.

```go
    type A struct {
        IntField int
    }

    p := new(A) // same as &A{}
```

Pointers behave the same way as regular variables. You can copy them by assigning a pointer type to other variables, pass and return them to functions, and create pointers to them.

The type of pointer to a pointer is described as `**T`, for example `**int`.

To get or change the value stored by the pointer, use the **dereference operator** `*`.

```go
i := 42
p := &i
fmt.Println(*p)
*p = 21
```

Calling the dereference operator on a `nil` pointer will lead to a panic at the code execution stage, and the program will refuse to continue working.

```go
var p *int
fmt.Println(*p) // panic: runtime error: invalid memory address or nil pointer dereference
```

### Pointers and structures

For pointers to structures in Go, there is the possibility of implicit dereference when accessing the fields of the structure.

```go
type A struct {
    IntField int
}

p := &A{}
p.IntField = 42 // instead of (*p).IntField or p->IntField like in C
```

### Comparing pointers

Comparison operators `(==, !=)` are defined for pointers. Two pointers are equal if they point to the same object in memory or if both are equal to `nil`.

#### When to use pointers

- When you need to change the value of a variable from a called function. If you pass a variable by value, all modifications inside the function will be applied to the local copy and leave the original variable unchanged.
``` go
  incrementCopy := func(i int) {
      i++
  }

  increment := func(i *int) {
      (*i)++
  }

  i := 42

  incrementCopy(i)
  fmt.Println(i) // 42

  increment(&i)
  fmt.Println(i) // 43

```

- When you need to emphasize that the value may be missing. For example, there is a function that returns a record about a user `type User struct{...}` by his ID. The index result makes it clear that not all IDs can be used to find the user. An example of a function with this signature:

```go
  func FindUser(id int) *User
```

- When you're working with resources like file descriptors or sockets. Copying of such variables may be due to the exhaustion of system resources or may not be performed at all.

- When you're working with large variables and copying across the stack takes more resources than garbage collection from pointers.

#### When not to use pointers

- When you want to speed up an application and it seems that copying structures is too expensive an operation. As long as there are no tests that clearly show that pointers improve performance, it is better not to try to optimize. Most likely, you will waste your energy or reduce system performance by increasing the cost of garbage collection.

- It is worth thinking about replacing value transfer with pointer transfer when the size of the structure reaches the order of hundreds of bytes.

- When there are a lot of pointers in memory, the garbage collector is heavily loaded. This can happen, for example, when creating your own in-memory database.

### Comparison of pointers in Go and C/C++

The syntax of pointers in Go is identical to their C counterparts, as are many other parameters. However, there are a couple of important differences.

Pointers in Go do not have address arithmetic. Despite the fact that the pointer stores the address, which is a number, arithmetic operations cannot be applied to it. This is not a disadvantage, because it was deliberately removed to increase the security of the code.

A pointer may not "look" at any part of memory, but only at an existing and corresponding type of pointer.

In fact, the pointer is close to immutability — you can create it and assign addresses to existing variables. In C, this is implemented much more broadly.

The garbage collector, one of the key features of Go, will not be able to delete a variable while any pointer is "looking" at it. Therefore, you can do without manually releasing memory using the `free` operation.

### Example

Let's imagine some structure that describes the user:

```go
type Person struct {
  Name string
  Age int
  lastVisited time.Time
}
```

In the `lastVisited` field, you need to save the date of the last visit. If the `Person` type is used in another package, then `lastVisited` cannot be changed directly, because the field is not exported. Without pointers, the function would look something like this:

```go
func GetPersonWithLastVisited(p Person) Person {
  p.lastVisited = time.Now()
    return p
}

// usage in other package
p := Person{
  Name: "Alex",
  Age: 25,
}
p = GetPersonWithLastVisited(p)
```

With this approach, you should not forget to assign a return value to the variable, since a copy of the object is passed in the parameter and all its changes inside the function do not affect the original variable. Also, if the argument type has a large size, then creating a copy may take longer than transferring the object to the address.

With pointers everything is easier:

```go
func UpdatePersonWithLastVisited (p *Person )  {
  p.lastVisited = time.Now()
}

// usage in other package
p := Person{
  Name: "Alex",
  Age: 25,
}
UpdatePersonWithLastVisited(&p)
```

Here, a pointer to a variable was passed to the function, which made it possible to change its field without additional copying. On the other hand, if a function does not change the parameters passed to it, then it is better not to use pointers, but to pass variables by value.
