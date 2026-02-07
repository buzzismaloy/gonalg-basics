# Fourth lesson

## Structures

A structure in Go is a data type with a set of attributes (fields) used to describe composite objects. The structure has similar analogues in other programming languages:

    - C — `struct`;
    - C++ — `class`, `struct`;
    - Python — `class`, `tuple`;
    - PHP — `class`;
    Lua — `table`.

See what the description of the `Person` type looks like:

```go
type Person struct {
    Name        string
    Email       string
    dateOfBirth time.Time
}
```

The structure fields can be of any type available in the language.

They can also be pointers to the structure itself. A classic example is the tree data structure:

```go
type Tree struct {
    Value      int
    LeftChild  *Tree
    RightChild *Tree
}
```

### Initialisation

There are several approaches to creating an instance of an object.

#### 1. Empty object

```go
p := Person{}
// or
var p Person
```

With this approach, all fields of the structure take default values.
The approach is applied:

    - when an instance does not require special initialization and can be used further in the code;
    - when additional conditions and data are needed to initialize fields, that is, setting the values of specific fields will follow the code below.

#### 2. Implicit indication of field values

```go
date := time.Date(2000, 12, 1, 0, 0, 0, 0, time.UTC)
p := Person{"Ivan", "Ivan@therealexistingmail.ru", date}
```

With this approach, the values for all fields of the structure are listed using literals or values of external variables.
Requirements:

    - You need to list all the fields of the object.
    - The order of the initializer arguments must match the order of the description of the structure fields. If you put the `Email` field in the first place in the description of the `type Person struct`, initializing the instance above will be incorrect (from the point of view of logic, but not the compiler).

The approach is applied:

    - when you need to explicitly specify the values of all fields of an object;
    - when you are sure that the type specification will not change often, otherwise you will have to make edits for each initializer of the object in the code.

#### 3. Explicit indication of field values

```go
p := Person{Name: "Ivan", Email: "Ivan@therealmail.ru"
```

With this approach, the names of the fields and their values are explicitly indicated.
Features:

    - this approach differs from the first by optionally specifying fields;
    - the order in which the fields are specified is not important;
    - the values of the fields that were not used in the initializer (`dateOfBirth` in the example) will take the default values.

To increase the readability of the code, such initialization is often described in several lines, which is also true for the second approach:

```go
p := Person{
    Name:  "Ivan",
    Email: "Ivan@realmail.ru",
}
```

***Note that the last line in the multiline literal entry also ends with a comma. This is done so that you can insert and delete lines without worrying about the comma at the end.***

The approach is applied:

    - almost always, as it lacks the limitations described above.

In practice, explicit naming is usually used because it reduces the number of possible errors.
Let's look at an example where a `Person` structure with the following fields is declared in the code, possibly in another package or even a repository:

```go
type Person struct {
    Name string
    Age int
}
```

The structure is very simple, and Go initialises it as follows:

```go
man := Person{"Alex", 30}

fmt.Printf("Man %#v", man)
```

Then the description of the structure is changed:

```go
type Person struct {
    Name     string
    NumChild int
    Age      int
}
```

The compiler suggests that there are not enough fields for initialization, then we add another value:

```go
man := Person{"Alex", 30, 2}

// and we get a program that works even though it doesn't work in a proper way
fmt.Printf("Man %#v", man)
```

If the fields were explicitly specified, there would be no problem:

```go
man := Person{Name: "Alex", Age: 30}
```

#### 4. Constructor

Given the subtleties of initializing a complex object, developers use constructors.

There is no syntax for constructors and destructors in Go, but you can often find an analog:

```go
   func NewPerson(name, email string, dobYear, dobMonth, dobDay int) Person {
       return Person{
           Name:        name,
           Email:       email,
           dateOfBirth: time.Date(dobYear, time.Month(dobMonth), dobDay, 0, 0, 0, 0, time.UTC),
       }
   }
```

Here are some rules approved by the Go community:

    - the name of the constructor function is written with the prefix `New`;
    - if the constructor validates arguments, the function must return an error with the last argument.

We can go back to our example to add a verification of the correctness of the `email` and the numeric components of the date, then the function declaration will take the form:

```go
func NewPerson(name, email string, dobYear, dobMonth, dobDay int) (Person, error) {}
```

The approach is applied:

    - when it is necessary to validate arguments in order to build a logically correct object;
    - when building an instance of an object requires additional actions, for example, connecting to a database.

***The example above is not ideal, since changing the `Person` specification will require changing the constructor prototype or creating a new version (say, `NewPersonWithPhone()`).***

### Access to fields

A dot is used to access the fields of the structure (`p.Name`):

```go
p := NewPerson("Ivan", "Ivan@realmail.ru", 2000, 12, 1)
fmt.Println(p.Name, p.Email)

p.Name = "Vasya"
fmt.Println(p.Name)
```

output:

```
Ivan Ivan@realmail.ru
Vasya
```

### Scope

As you already know, Go has concepts of **exported** and **non-exported types**. The code is divided into packages, and in order for a type, function, or global variable to be available in another package, their names must begin with a capital letter. The same rule applies to fields and methods of the structure.

In the example above, `Person` is the exported type (public). Other packages can create instances of this type and have access to the public `Name` and `Email` fields. And the `dateOfBirth` field is non—exportable (private).

***The non-exported type can be used in another package if there is a corresponding type constructor and an exported constructor function. This trick is found in Go code, but interfaces are most often used to hide the implementation.***

An example of exporting a private type is given in `example-project/`

### Tags

Each field of the structure can have a set of annotations, which are called tags:

```go
type GetUserRequest struct {
    UserId string `json:"user_id" yaml: "user_id" format:"uuid" example:"2e263a90-b74b-11eb-8529-0242ac130003"`
    IsDeleted *bool `json:"is_deleted,omitempty" yaml:"is_deleted"`
}
```

Tags do not affect the presentation or working with data directly, but can be used by packages to get additional information about a specific field.

A set of tags with their values can be represented as a set of keys and values, where the keys are separated by spaces and the key values are separated by a comma.

In the example above, the following tags are found:

    - `json` — used by the [encoding/json](https://golang.org/pkg/encoding/json/) package to serialize/deserialize structures in JSON;
    - `yaml` is similar to `json`, but is used by external libraries to work with the YAML format.;
    - `format` and `example` can be both a hint for the developer, and an annotation for generating a Swagger description (for example, the [swag](https://github.com/swaggo/swag) library).

The annotations used most often depend on the library used. Possible keys and values should be searched in the package documentation (in the worst case, in the code).

The developer can enter their tags and work with them through the `reflect` package of the standard library.

## Anonymous structures

Anonymous structures are declared and used directly in the code. A separate type is not described for them, because anonymous structures are used only once, and the description makes sense only for a specific part of the code: for example, when serializing/deserializing messages. Anonymous structures are most often used in tests to describe test structures.

The easiest way to understand the concept of anonymous structures is from the following consideration:

```go
type Person struct {
    Name string
}
```

The `type Person ...` construction in fact, it does not describe, but creates a type based on an existing one and names it. That is, in fact, it is the `struct{}` construction that creates the type.

Having received such an anonymous type, you can immediately create a variable of this type.
Here is an example of using an anonymous structure when building a REST query:

```go
req := struct {
    NameContains string `json:"name_contains"`
    Offset       int    `json:"offset"`
    Limit        int    `json:"limit"`
}{
    NameContains: "Иван",
    Limit:        50,
}

reqRaw, _ := json.Marshal(req)
fmt.Println(string(reqRaw))
```

output:

```
{"name_contains":"Иван","offset":0,"limit":50}
```

Here we described an anonymous structure, initialized its instance, performed JSON serialization, and output the result as a string.

### struct{}

```go
var c struct{}
// or
c := struct{}{}

fmt.Println(unsafe.Sizeof(c))
fmt.Println(unsafe.Pointer(&c))
```

output:

```
0
0x11d46e8
```

The size of struct{} is 0, and object c has an address. This loophole can be used to optimize the code from memory, and in the future we will analyze this in practice.
