# First lesson

## OOP

In this topic, you'll expand your understanding of the capabilities of structures in Go:

- You'll become familiar with the syntax for describing their methods;
- You'll learn how to "nest" one structure within another;
- You'll discover whether Go is an object-oriented programming language.

The foundation of OOP consists of the following components:

- **Abstraction** — the ability to define properties and methods of an object that fully describe its characteristics, define the boundaries of its application, and allow the object to be used as an integral part of a system.
- **Encapsulation** — the ability to tie an object's data and its behavior together.
- **Hiding** — the ability to hide an object's implementation by providing the user with a specification (interface) for interacting with it (hiding something "under the hood").
- **Inheritance** — the ability to create derivatives of parent objects that extend or modify the properties and behavior of the parent (base class).
- **Polymorphism** — the ability to implement objects with the same specification differently.

Let's consider how each of these OOP mechanisms is implemented in Go and whether they even exist.

### Methods

A method is a function bound to a specific type. Methods allow a type's behavior and data to be bound within the type itself, providing encapsulation. If you're familiar with Python or C++/C#/Java programming, methods in Go will be similar to class methods, with a few exceptions.

***In Python, for example, a class object is passed to a method through the explicitly declared first variable, self. Go has a separate syntax for this: the method must be declared in the same package as the type it is bound to. In Python, the method is defined within the class body. C++ requires explicitly specifying the class to which the method belongs, but a pointer to the class object for which the method is called is passed implicitly through the this variable.***

```go
package main

import "fmt"

 // type declaration
type MyType int

// method declaration
func (m MyType) String() string{
    return fmt.Sprintf("MyType: %d", m)
}

func main() {
    var m MyType = 5

    // method call
    s := m.String()
    fmt.Println(s)
}
```

The syntax of a type method is similar to that of a regular function, but a **receiver** is added after the `func` keyword. You could say that the receiver is another argument to the function.
Methods aren't limited to the struct data type, which may surprise those who have previously written in other OOP languages. Since methods are defined in the same package as the type, all unexported elements in that package are accessible to them.
It's important to note that methods don't inherently provide a method hiding mechanism. If we define two types in the same package, the unexported fields and methods of one type will be accessible to methods of the other type. Hiding is achieved through the unexported elements of the package.

Here's an example of a canonical Go enum implementation with a minimal set of methods:

```go
// DeliveryState — the status of message delivery and processing.
type DeliveryState string

// Possible values of the DeliveryState enum.
const (
    DeliveryStatePending DeliveryState = "pending" // message sent
    DeliveryStateAck DeliveryState = "acknowledged" // message received
    DeliveryStateProcessed DeliveryState = "processed" // message processed successfully
    DeliveryStateCanceled DeliveryState = "canceled" // message processing interrupted
)

// IsValid checks the validity of the current value of DeliveryState type.
func (s DeliveryState) IsValid() bool {
    switch s {
    case DeliveryStatePending, DeliveryStateAck, DeliveryStateProcessed, DeliveryStateCanceled:
        return true
    default:
        return false
    }
}

// String returns a string representation of the DeliveryState type.
func (s DeliveryState) String() string {
    return string(s)
}
```

The `DeliveryState` type is equivalent to the `string` type, so you can obtain an instance of it with a simple type cast. Here's an example of such a cast and the use of a validation function:

```go
func HandleMsgDeliveryStatus(status DeliveryState) error {
// Checking the validity of an enum value by calling a method on the DeliveryState type
    if !status.IsValid() {
        return fmt.Error("status: invalid")
    }

// Message processing code

    return nil
}

func main() {
// Cast the string "fake" to the DeliveryState type
    if err := HandleMsgDeliveryStatus(DeliveryState("fake")); err != nil {
        panic(err)
    }
}
```

### Methods of structures

Let's look at another example of using methods. Methods are defined for user-defined data types, and most often, structures serve this role. In general, structure methods are not clearly different from methods of other types, but there are some nuances worth exploring.

Let's define a `circular buffer` type with a minimal set of methods and provide a code example:

```go
// CircularBuffer implements a "circular buffer" data structure for float64 values.
type CircularBuffer struct {
    values []float64 // current buffer values
    headIdx int // head index (first non-empty element)
    tailIdx int // tail index (first empty element)
}

// GetCurrentSize returns the current buffer length.
func (b CircularBuffer) GetCurrentSize() int {
    if b.tailIdx < b.headIdx {
        return b.tailIdx + cap(b.values) - b.headIdx
    }

    return b.tailIdx - b.headIdx
}

// GetValues returns a slice of the current buffer values, preserving the order of writes.
func (b CircularBuffer) GetValues() (retValues []float64) {
    for i := b.headIdx; i != b.tailIdx; i = (i + 1) % cap(b.values) {
        retValues = append(retValues, b.values[i])
    }

    return
}

// AddValue adds a new value to the buffer.
func (b *CircularBuffer) AddValue(v float64) {
    b.values[b.tailIdx] = v
    b.tailIdx = (b.tailIdx + 1) % cap(b.values)
    if b.tailIdx == b.headIdx {
        b.headIdx = (b.headIdx + 1) % cap(b.values)
    }
}

// NewCircularBuffer - constructor of the CircularBuffer type.
func NewCircularBuffer(size int) CircularBuffer {
    return CircularBuffer{values: make([]float64, size+1)}
}

func main() {
    buf := NewCircularBuffer(4)
    for i := 0; i < 6; i++ {
        if i > 0 {
            buf.AddValue(float64(i))
        }
        fmt.Printf("[%d]: %v\n", buf.GetCurrentSize(), buf.GetValues())
    }
}
```

The example above shows that a method can have two receiver options:
1. Receiver by value (`b CircularBuffer`).
2. Receiver by pointer (`b* CircularBuffer`).

#### Receiver by value

***Note***: The call is the same in both cases. However, with `b.Method()` for pointer receivers, the compiler actually generates a call to `(&b).Method()`, meaning the method is passed a pointer to the object on which the method is called.

The `CircularBuffer` type has two methods with a **value receiver**: `GetCurrentSize` and `GetValues`. For these methods, the receiver takes the form `func(b CircularBuffer)`.

Both methods do not modify the state of the object from a type logic perspective. From a language perspective, methods with such a receiver cannot modify the state of the object that called the method.

The variable `b` (the method receiver) contains a copy of the `CircularBuffer` instance, so any change to `b`'s fields will modify the local (for the method) copy of the object. If a structure field is a reference type, changing this field in the local copy will affect the original variable. For example, let's add a method that explicitly sets the value of a buffer element (the `values` slice) by index:

```go
// ForceSetValueByIdx sets the buffer value by index.
func (b CircularBuffer) ForceSetValueByIdx(idx int, v float64) {
// It's best not to use this technique in practice when the method parameter
// is not a pointer, but a value
    b.values[idx] = v
}

func main() {
    buf := NewCircularBuffer(4)
    buf.ForceSetValueByIdx(0, -1.0)
    buf.ForceSetValueByIdx(1, -2.0)
    fmt.Println(buf.values)
}
```

output:

```
[-1 -2 0 0 0]
```

Why does this happen? If a field is a pointer or has a reference type (`map, chan, slice`), it will reference the same objects in the copy of the variable.

#### Receiver by pointer

The `CircularBuffer` type has one method with a **pointer receiver**: `AddValue`. This type's receiver takes the form `func(b *CircularBuffer)`. The method function receives a pointer to the type instance and, as a result, can modify its fields.

It's important to note that pointer receiver methods can be found not only in structs, but also in any other user-defined type:

```go
type IntSlice []int

func (s *IntSlice) Add(v int) {
    *s = append(*s, v)
}

func main() {
    s := make(IntSlice, 0)
    s.Add(1)
    s.Add(2)
    fmt.Println(s)
}
```

output:

```
[1 2]
```

A method with a by-value receiver receives a copy of the object on which it was called, so such a method cannot change the value of the calling object. However, if the object's fields contain reference types, their changes will affect the original object. A method with a by-pointer receiver receives a pointer to the object on which it was called and operates on that pointer.

- Calling a by-value method on a pointer to an object is equivalent to calling `(*b).Method()`.
- Calling a by-pointer method on a value object is equivalent to calling `(&b).Method()`.

### OOP and Methods

As shown above, methods are responsible for encapsulation in Go's implementation of the OOP paradigm and truly allow behavior and data to be bound together in a single type.

***Note***: since a function in Go is a first-order object, this provides several opportunities for expanding programming capabilities and producing more elegant code.

#### Function as a field of a structure

For example define a following structure:

```go
package main

import "fmt"

type MyStruct struct {
    A   int
    Log func(s string)
}

func main() {
    var s = MyStruct{
        A:   1,
        Log: func(s string) { fmt.Println(s) },
    }

    s.Log("some string")
}
```

Calling a field function is externally no different from calling a method, but there are some key differences:

1. A field function does not have access to the object that called it unless it is explicitly passed to it.
2. A field function can be dynamically redefined at runtime. This allows, for example, the use of functions from other packages.
3. A field function can be empty. In this case, calling it will cause a panic.

A function as a structure field can be used to change the behavior of an object on the fly.

#### Passing a method as a function argument

We already know that a function in Go can be passed as an argument to another function. Methods work similarly.
Let's look at an example. We have a handler function that accepts a number and a function:

```go
func Handle(num float64,  add func(float64)) {
  add(num)
}

// Let's create a new ring buffer from the example above:
buf := NewCircularBuffer(4)

// Let's call
Handle(1.0, buf.AddValue)
Handle(2.0, buf.AddValue)
Handle(3.0, buf.AddValue)
Handle(4.0, buf.AddValue)
fmt.Printf("[%d]: %v\n", buf.GetCurrentSize(), buf.GetValues())
```

This demonstrates a very important point: the method was passed as a function to the handler function. However, it retains its binding to the specific instance of the structure of which it is a method.

The **handler argument type** is the function type, and the receiver of this function can be any type. That is, the methods' types match if their arguments and return values match. The receiver type is not taken into account. This approach is often used when building servers, where methods are registered as handlers for incoming requests.

Note that the `CircularBuffer` type has methods with both value and pointer receivers. This mixing of different method types is permitted by the language standard, but is not common in the Go community. Adhere to the conventions your development team follows.
If an object has only value-based methods and all fields are non-exported, it can be said to be **immutable**. Conversely, an object is **mutable** if all methods have pointer receivers or one of the fields is exported. This makes life a lot easier for your type of user if you're developing a library with tens of thousands of stars on GitHub.

Having methods on a pointer does not force you to create instances of it through the pointer:

```go
type MyType struct {
    value int
}

func (t *MyType) SetValue(v int) {
    t.value = v
}

func (t MyType) String() string {
    return fmt.Sprintf("Value: %d", t.value)
}

func main() {
    t := MyType{}
    // or
    t := &MyType{}

    t.SetValue(100)
    fmt.Println(t)
}
```

output:

```
Value: 100
```

Thus, methods with a pointer receiver and a value receiver work almost identically, except for the difference between what is passed to the method function—a value or a pointer.
Methods are one of the fundamental and, arguably, the most widely used tools in Go for building complex programs.
Methods are not class members, as in Python, and can be created for any data type. Variables are accessed from within a method through the receiver.
