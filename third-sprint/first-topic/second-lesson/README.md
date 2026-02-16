# Second lesson

## Embedding

Inheritance is one of the fundamental principles of OOP and describes the relationship between classes based on who/what they represent. For example, a duck is a bird. A bird is an animal, meaning the inheritance hierarchy is built toward the most abstract object.
One could say that inheritance should imply that all ducks are birds. However, the converse is not true, since not all birds are ducks.

However, Go uses composition instead of inheritance. Structures can contain other structures and types.

**Composition** is a relationship that contains something. For example, a duck consists of a beak, a body, and feet. The beak allows it to quack, so a duck can quack. Not everything with a beak is a duck. But all ducks have a beak.
Embedding, or inlining, is the implementation of composition in Go.

In general, embedding looks like this:

```go
type OuterStruct struct {
    EmbeddedType
    A int
    B int
}
```

All fields and methods of the `EmbeddedType` will be passed to the `OuterStruct` struct as if it contained them. This allows for code reuse in complex structures by embedding them within each other.
Let's look at an example. Let's create two structures, `Person` and `Student`. `Student` will contain `Person`. This is similar to inheritance, but there are significant differences. The `Student` struct is not a `Person`; rather, it contains it.

A `Student` object cannot be cast to the `Person` type using type casting. This refers to the classic type casting in OOP languages, where an instance of a derived class can act as an instance of a base class.
For example, there is no such construct as `person := Student(person)` for built-in types in Go.

```go
package main

import (
    "fmt"
)

// Person — structure describaing a human being.
type Person struct {
    Name string
    Year int
}

// NewPerson returns new Person struct.
func NewPerson(name string, year int) Person {
    return Person{
        Name: name,
        Year: year,
    }
}

// String returns name of the Person.
func (p Person) String() string {
    return fmt.Sprintf("Имя: %s, Год рождения: %d", p.Name, p.Year)
}

// Print prints Person's info.
func (p Person) Print() {
    // call of the String method of for Person struct
    fmt.Println(p)
}

// Student structure describes a student with embedded Person structure
type Student struct {
    Person // embedded object Person
    Group  string
}

func NewStudent(name string, year int, group string) Student {
    return Student{
        Person: NewPerson(name, year), // Create structure Person
        Group:  group,
    }
}

// String returns info about student
func (s Student) String() string {
    return fmt.Sprintf("%s, Группа: %s", s.Person, s.Group)
}

func main() {
    s := NewStudent("John Doe", 1980, "701")
    s.Print()
    // the String() method of Student is called
    fmt.Println(s)
    fmt.Println(s.Name, s.Year, s.Group)
}
```

output:

```
Имя: John Doe, Год рождения: 1980
Имя: John Doe, Год рождения: 1980, Группа: 701
John Doe 1980 701
```

In this example, the `Student` type inherits all the fields and methods of the `Person` object. Since the `Print` method is defined only for the `Person` type, it only displays the name and year of birth. When calling the `fmt.Println` function, the `String()` method, which is defined for the `Student` type, is used, so all student information is displayed.

### Accessing Nested Structure Fields

If a nested type is defined in another package, the type using it has access only to exported (public) methods and fields. There are several ways to provide access to nested structure fields. Let's demonstrate them by adding a `Debug()` method to the `Student` type:

```go
func (s *Student) Debug() {
    // access to Person's methods
    s.Print()
    // or
    s.Person.Print()

    // access to Name of Person
    s.Name = "Mark Smith"
    // или
    s.Person.Name = "Mark Smith"

    // the method String for Student is called
    fmt.Println(s)
    // the method String for Person is called
    fmt.Println(s.Person)
}
```

A method of a structure that contains another type is not overridden, but shadowed. This means that when calling a method with the same name, Go first tries to find the method among the methods of the `Student` structure, then searches for it among the nested types. If multiple nested types have the same methods, a conflict occurs.

Fields (or methods) of a nested structure can be accessed either with or without specifying the object type (`s.Person.Print()`). The `Person` object's `String` method requires explicitly specifying the type, since the method is overridden by the `Student` type.

In fact, the structure has a field with the same name as the type. However, when calling a method of a nested structure, it is not the variable that called it that is passed into it, but this field. If necessary, you can explicitly access the method of the nested structure.

There are two rules for resolving field and method name conflicts:

1. A named field (method) of a structure shadows a field (method) of the same name for nested structures. The top-level name dominates names of lower levels. For example, calling the `student.String()` method calls a method of the `Student` structure, not the `Person` structure.
2. If a field (method) name appears at the same nesting level (name duplication) and is used in the code, this is an error. If name duplication exists but the name is not used in the code, the compiler will not generate an error. For example, if the `Student` type had another nested structure, `Faculty`, with a `Print()` method, then the nested type name `s.Person.Print() or s.Faculty.Print()` would need to be specified when calling the method.
For example:

```go
    type Faculty string

    func (f Faculty) String() string {
        return string(f)
    }

    func (f Faculty) Print() {
        fmt.Println(f)
    }


    type Student struct {
        Person
        Faculty
        Group  string
    }

    func main {
    //    s.Print()
        fmt.Println(s)
    }
```

If we uncomment the `s.Print()` line in the example above, we'll get a compilation error. The compiler can't determine the `ambiguous selector s.Print` precisely, and tries to figure out which method to select from `Faculty or Person`. This results in an error.

However, the `String()` method, which is called in `Println()`, hides all `String()` methods of nested types. This prevents the error from occurring.
It's worth noting that these same rules apply not only to methods but also to structure fields.

You can rewrite the example without using nested structures and work with the required objects as with regular fields:

```go
type Student struct {
    Person Person
    Group string
}
```

In this case, the `Person` type will no longer be nested, and you must explicitly specify `s.Person.Print(), s.Person.Name`, etc.
The embedding mechanism is separate from Go interfaces, which we'll discuss in more detail in the next topic. It's perceived as syntactic sugar, but it becomes useful when you need an object to conform to specific interfaces. If a type has a nested type, it implements all the interfaces of that type. You can read more about syntactic sugar [here](thecode.media/sugar).

### Embedding Type Pointers

A structure can embed a type pointer. Let's look at an example:

```go
    type Student struct {
        *Person
        Group  string
    }
```

The external differences are minor, but pointer embedding can be convenient if you're embedding a larger structure into your struct and then passing it by value. This can improve performance because only the pointer to the struct is copied, not the entire struct. It's important to note that in this case, the embedded struct can be modified:

```go
    // gets structure by value
    func ChangeName( s Student, name string) {
        s.Name = name
    }

    s := Student{&Person{Name: "alex"}, "021"}
    ChangeName(s, "teodor")
    fmt.Println(s.Name) // "teodor"
```

This function will change the student's name if the `Person` field is embedded as a pointer. If you simply embed `Person`, the name will remain unchanged:

```go
    type Student struct {
        Person
        Group  string
    }

    func ChangeName( s Student, name string) {
        s.Name = name
    }

    s := Student{Person{Name: "alex"}, "021"}
    ChangeName(s, "teodor")
    fmt.Println(s.Name) // "alex"
```

### Restrictions on Embedded Types

According to the Go specification, not all data types can be embedded in structs. Types or pointers to types can be embedded. However, if a pointer to a type is embedded, the type itself cannot be a pointer. This means that, for example, you can embed the `Person` type and the `*Person` type, but you cannot embed the `**Person` type.

#### Usage

Another use for embedding is extending the capabilities of external types. You might want to have all the capabilities of a type in an external package, but you can't modify the package itself. In this case, you can create your own type, embed the external type, and add the necessary methods and fields. The resulting type will contain the methods of the embedded type.
Library developers often assume that the type provided by the library will be embedded into the user's structure, rather than used independently.

Let's take the `CircularBuffer` type from the previous lesson and add a method that can add multiple values to our buffer at once. Let's imagine that its code can't be changed (this is very common in production development):

```go
// ExtendedCircularBuffer — «inheritor» of CircularBuffer type, that implements adding of multiple elements
type ExtendedCircularBuffer struct {
    CircularBuffer
}

func (cb * ExtendedCircularBuffer) AddValues(vals... float64)  {
    for _,val := range vals {
        cb.Addvalues(val)
    }
}

func NewExtendedCircularBuffer (size int) ExtendedCircularBuffer {
    return ExtendedCircularBuffer{
        CircularBuffer : NewCircularBuffer(size),
    }
}


func main() {
    buffer := NewExtendedCircularBuffer(5)
    buffer.Addvalues(1,2,3,4,5)
    fmt.Printf("[%d]: %v\n", buffer.GetCurrentSize(), buffer.GetValues())
}
```

output:

```
[5]: [1 2 3 4 5]
```

P.S. Educational project(extended logger) is [here](https://github.com/buzzismaloy/go-logextended).
