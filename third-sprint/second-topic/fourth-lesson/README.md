# Fourth lesson

## Reflection

### What is reflection

**Reflection** in programming refers to the ability to obtain type information from a variable of that type. Simply put, reflection allows you to obtain information about program code and modify it at runtime.
Reflection is typically used to work with data whose type is unknown at compile time. For example, data may arrive over the network that needs to be stored in structures. However, you may not know what this data is, so in such cases, the ability to create data structures on the fly is essential.

It's worth noting that reflection isn't used in most developer tasks. However, it is used in the popular `encoding/json` package and many others. This is a good reason to learn reflection and understand how it works.

***Note***: To work with reflection in the Go language, there is a [reflect](https://golang.org/pkg/reflect/) package from the standard library.

In this lesson, you'll learn about some of the capabilities of the `reflect` package. The package's functions work with arbitrary static types (`interface{}`) and allow you to obtain metainformation about them. Using this package, you can dynamically create types at runtime.

#### DeepEqual

Sometimes you need to compare two variables of the same type by value, and sometimes the simple approach of using `==` doesn't work. Then you need to look deeper, comparing all the values stored in slices and maps, under pointers.
Let's look at a simple example of a naive approach:

```go
type MyType struct {
    IntField   int
    StrField   string
    PtrField   *float64
}

func (mt MyType) IsEqual(mt2 MyType) bool {
    return mt == mt2
}

func main() {
    floatValue1, floatValue2 := 10.0, 10.0
    a := MyType{IntField: 1, StrField: "str", PtrField: &floatValue1}
    b := MyType{IntField: 1, StrField: "str", PtrField: &floatValue2}

    fmt.Printf(" a == b: %v\n", a.IsEqual(b))
}
```

output:

```
a == b: false
```

We get `false`, even though the contents of these variables appear equal. What's the catch?
The example defines a `MyType` type and a method for comparing two instances of the `IsEqual` type.
Running this snippet prints `false` to the console, even though the values of all fields of objects `a` and `b` are equal. This happens because direct pointer comparison (the `PtrField` field) compares addresses, not values. If both pointers are `nil`, the code will print true.

Let's change the type specification:

```go
type MyType struct {
    IntField   int
    StrField   string
    PtrField   *float64
    SliceField []int
}
```

Compiling this code will generate an error because the `==` operator is not defined for the `SliceField` type. Adding a reference type field to a structure (or to a field in a nested structure) means that the `==` operator cannot be used for direct comparison. This is true not only for structures, but also for all user-defined types, such as `type MySlice []int`.

***One solution is to modify the `IsEqual` method and add several `if` statements. Writing a comparison method or function is common practice in Go, as the language doesn't allow operator overloading (`==` in this case).***

The `reflect` package offers the following solution:

```go
func (mt MyType) IsEqual(mt2 MyType) bool {
    return reflect.DeepEqual(mt, mt2)
}

func main() {
    floatValue1, floatValue2 := 10.0, 10.0
    a := MyType{IntField: 1, StrField: "str", PtrField: &floatValue1, SliceField: []int{1}}
    b := MyType{IntField: 1, StrField: "str", PtrField: &floatValue2, SliceField: []int{1}}

    fmt.Printf("a == b: %v\n", a.IsEqual(b))
}
```

output:

```
a == b: true
```

The function compares the values of all elements of a type, including nested elements. All comparison criteria for the `DeepEqual` function can be found in the [documentation](https://golang.org/pkg/reflect/#DeepEqual).
In practice, `DeepEqual` is rarely used because calling this function recursively traverses all elements of a type, which consumes a lot of CPU time. `IsEqual` is most often written manually, limiting the scope of the comparison to the required application logic.
`DeepEqual` allows you to compare two variables of the same type by value, even if these variables have a complex data structure, such as containing references to other variables.

#### Value and ValueOf()

In the previous lesson, we discussed interface casting. We can cast an interface of one type to another. However, what if we don't know what type the passed object should be cast to?

The `reflect` package provides advanced capabilities for working with type casts and inspecting them at runtime.
Imagine we're writing a library that needs to work with a variety of data types. An example would be the `encoding/json` package, which can accept and serialize any structure.

Casting alone is clearly not enough here, as developers often face questions like:

- How many fields does the structure have?
- What is their type?
- What are their names? We want to get these names as a string.

At the same time, the structure is passed to our library wrapped in an empty interface.
These are precisely the questions that reflection is designed to answer.

Every value, regardless of type, can be cast to the universal type `reflect.Value`. This is done by calling `reflect.ValueOf(v inteface()) Value`. This function takes a value and returns `Value`. `Value` itself has many methods that allow you to obtain information about both the type and the value.
Let's look at some of them.

##### Type & kind

The `Type()` method returns the object's type.
The `Kind()` method returns the object's base type, which is not a user-defined type, but one of the built-in types in the Go language: `structure, channel, slice, function, array, and others`.

```go
var varBool *bool
fmt.Println(reflect.ValueOf(varBool).Kind()) // ptr — pointer
fmt.Println(reflect.ValueOf(varBool).Type()) // *bool — bool pointer

var varFloat float32
fmt.Println(reflect.ValueOf(varFloat).Kind()) // float32
fmt.Println(reflect.ValueOf(varFloat).Type()) // float32

var varMap map[string]int
fmt.Println(reflect.ValueOf(varMap).Kind()) // map
fmt.Println(reflect.ValueOf(varMap).Type()) // map[string]int

varStruct := struct{Value int}{}
fmt.Println(reflect.ValueOf(varStruct).Kind()) // struct
fmt.Println(reflect.ValueOf(varStruct).Type()) // struct { Value int }
```

The `reflect.Type` type describes a Go type. The `reflect.Kind` type defines many basic Go types: struct, channel, slice, function, array, and others. In other words, `Type` describes the specific type of a value, and `Kind` describes what kind of type it is.
This is quite useful for understanding whether we're being passed a structure, an array, or just an integer.

##### Checking for nil

In the previous lesson, we discussed comparing interfaces with `nil` and learned that `nil` can be either the value of the interface itself or the value of the value it points to.

Let's try implementing a naive implementation of this comparison.

```go
type MyType struct{}

func NaiveIsNil(obj interface{}) bool {
    return obj == nil
}

func main() {
    var t *MyType
    fmt.Printf("Checking the type (%v) for nil: %v\n", reflect.TypeOf(t), NaiveIsNil(t)) // TypeOf returns the type of transferred object.
}
```

output:

```
Checking the type (*main.MyType) for nil: false
```

Something clearly went wrong. We passed `nil` to the function and expected `true`.
The variable `t` is a pointer to the `MyType` type. When this variable is declared, the pointer doesn't specify a value, and technically it's empty, but a direct comparison with `nil` yields `false`. The reason for this was discussed in the previous lesson, so we won't go into detail here. Using the `reflect` package, we'll write a universal solution for this check:

```go
func IsNil(obj interface{}) bool {
    if obj == nil {
        return true
    }

    objValue := reflect.ValueOf(obj)
    // check the type is reference type so it may be equal to nil
    if objValue.Kind() != reflect.Ptr {
        return false
    }
    // check the value is equal to nil
    //  note: IsNil calls panic if value is not reference type so it is important to always check for a Kind()
    if  objValue.IsNil() {
        return true
    }

    return false
}
```

Here are a few more useful methods:

```go
varInt := 100
varIntValue := reflect.ValueOf(varInt)
fmt.Println(varIntValue.IsZero()) // false
fmt.Println(varIntValue.Int())    // 100

var varPtr *int
varPtrValue := reflect.ValueOf(varPtr)
fmt.Println(varPtrValue.IsNil())  // true
fmt.Println(varPtrValue.IsZero()) // true
```

The `IsZero()` method compares a value to the default value (0 for an `int`, `nil` for a pointer, etc.).
The `IsNil()` method compares a value to `nil` and is only applicable to types that support `nil` (`chan, slice, map, etc.`).
Like other types in the package, `Value` must be handled with care, as misuse of its methods will result in a panic. For example:

```go
var varBool *bool
varBoolValue := reflect.ValueOf(varBool)
fmt.Println(varBoolValue.IsNil())        // true
fmt.Println(varBoolValue.IsZero())       // true
fmt.Println(varBoolValue.Elem().Bool())  // panic: attempt to get a value for an empty Value
```

The `Elem()` method returns a value (also of type `Value`), which is described by the `Value` interface (`varBoolValue` in this case).
Essentially, the `Elem()` method is a dereference: it returns the value pointed to by `Value`, which was a pointer.
Let's correct the previous example and show how to assign a value to a pointer using reflection. We set the value of the `varBool` variable using the `Set(val reflect.Value)` method and passing it a `Value` of the appropriate type:

```go
var varBool *bool
fmt.Println(reflect.ValueOf(varBool).IsNil())  // true

trueVal := true
reflect.ValueOf(&varBool).Elem().Set(reflect.ValueOf(&trueVal))

fmt.Println(reflect.ValueOf(varBool).IsNil())       // false
fmt.Println(reflect.ValueOf(varBool).Elem().Bool()) // true — get value through reflection
fmt.Println(*varBool)                               // true — get value without reflection
```

#### Fields and NumFields — iterate over the fields of a structure

If you've ever written low-level C, you've probably worked with structures as byte arrays. You simply took a pointer to the structure and iterated through it, accessing its fields. This process is extremely dangerous and inconvenient, but it must be said that it's extremely effective in terms of performance.
Python also provides the ability to access class object elements as a dictionary. I'd like to see this capability in Go as well. The `reflect` package also provides this, but it's worth noting that this solution is slower than regular access to structure fields.

Here is an example:

```go
package main

import "reflect"

//
func ExtendedPrint(v interface{}) {
    val := reflect.ValueOf(v)
    // we check if we were passed a pointer to a structure
    switch val.Kind() {
    case reflect.Ptr:
        if val.Elem().Kind() != reflect.Struct {
            fmt.Printf("Pointer to %v : %v", val.Elem().Type(), val.Elem())
            return
        }
        // if this is a pointer to a structure, we will continue working with the structure itself
        val = val.Elem()

    case reflect.Struct: // work with structure
        fmt.Printf("Struct of type %v and number of fields %d:\n", val.Type(), val.NumField())
        for fieldIndex := 0; fieldIndex < val.NumField(); fieldIndex++ {
            field := val.Field(fieldIndex) // field — тоже Value
            fmt.Printf("\tField %v: %v - val :%v\n", val.Type().Field(fieldIndex).Name, field.Type(), field)
            // we get the field name not from the field value, but from its type.
        }
    default:
        fmt.Printf("%v : %v", val.Type(), val)
        return
    }
}

func main() {
    s := MyStruct{
        A: 3,
        B: "some",
        C: false,
    }
    s1 := &MyStruct{
        A: 7,
        B: "text",
        C: true,
    }

    ExtendedPrint(s)
    ExtendedPrint(s1)
    ExtendedPrint(struct {
        E int
        C string
    }{2, "other text"})
    ExtendedPrint("some string")
}
```

output:

```
Struct of type main.MyStruct and number of fields 3:
  Field A: int - val :3
  Field B: string - val :some
  Field C: bool - val :false
Struct of type main.MyStruct and number of fields 3:
  Field A: int - val :7
  Field B: string - val :text
  Field C: bool - val :true
Struct of type struct { E int; C string } and number of fields 2:
  Field E: int - val :2
  Field C: string - val :other text
string : some string

```

This provides a rather convenient and interesting feature that is often used for examining passed structures. If you need to access one of the fields of a passed structure by name, use these useful functions:

- `FieldByName(name string)` — returns the `Value` of a structure field by name.
- `FieldByIndex(i int)` — returns a structure field by index.

Similar functions exist for type methods.

#### Changing a structure field

Reflection can be used to modify a passed object.
However, not every object can be modified. To determine whether a `Value` can be modified, use the `CanSet()` method.
For example, let's implement a function that determines whether the input structure has a named field. If so, it modifies it.

```go
func ChangeFieldByName(v interface{}, fname string, newval int) {
    val := reflect.ValueOf(v)
    if val.Kind() == reflect.Ptr {
        val = val.Elem()
    }
    if val.Kind() != reflect.Struct {
        return
    }

    field := val.FieldByName(fname)
    if field.IsValid() {
        if field.CanSet() {
            switch field.Kind() {
            case reflect.Int:
                field.SetInt(int64(newval))
            case reflect.String:
                field.SetString(strconv.Itoa(newval))
            }
        }
    }
}
```

***Note***: `CanSet` checks whether changing the `Value` will affect the variable passed to it. If the variable is passed to `ValueOf` by value, the function will not work. Another important point: the value can only be changed for exported fields.

#### Dynamic type information (tag parsing)

For each structure field, you can specify tags — a string containing additional information about that field. For example, you can specify JSON serialization settings. Tags are specified in the key:value format, and multiple tags can be specified, separated by spaces. When parsing tags, they are converted into a set of `key-value` pairs.
Tags are typically used when information about a structure field obtained through reflection is insufficient. For example, you might want to know the exact name of a structure field for serialization.

Using the `reflect` package, let's write a `GetStructTags` function that will return information about all specified structure tags and their values:

```go
type (
    // FieldsInfo содержит информацию о полях структуры (ключ: имя поля).
    FieldsInfo map[string]FieldInfo

    // FieldInfo содержит информацию о поле структуры.
    FieldInfo struct {
        // тип поля
        Type     string     `json:"type"`
        // теги
        Tags     TagsInfo   `json:"tags,omitempty"`
        // информация по полям вложенной структуры
        Embedded FieldsInfo `json:"embedded,omitempty"`
    }

    // TagsInfo содержит информацию о тегах (ключ: имя тега).
    TagsInfo map[string][]string
)

// String возвращает строковую репрезентацию типа FieldsInfo.
func (f FieldsInfo) String() string {
    bz, _ := json.MarshalIndent(f, "", "   ")
    return string(bz)
}

// GetStructTags возвращает информацию по каждому полю структуры.
func GetStructTags(obj interface{}) (retInfos FieldsInfo) {
    retInfos = make(FieldsInfo)

    // получаем описание типа переданного объекта
    // далее по коду явно передаём в функцию тип `reflect.Type`, поддержим здесь этот случай рекурсивного вызова
    var objType reflect.Type
    if t, ok := obj.(reflect.Type); ok {
        objType = t
    } else {
        objType = reflect.ValueOf(obj).Type()
    }

    // чиним вход: если передали указатель, получим описание типа под указателем
    if objType.Kind() == reflect.Ptr {
        objType = objType.Elem()
    }

    // проверка входа: если объект не структура, искать теги не нужно
    if objType.Kind() != reflect.Struct {
        return
    }

    // итерируемся по всем полям структуры
    // NumField() — возвращает количество полей в структуре
    for fieldIdx := 0; fieldIdx < objType.NumField(); fieldIdx++ {
        field := objType.Field(fieldIdx) // получаем поле структуры
        retInfos[field.Name] = FieldInfo{
            Type:     field.Type.String(), // тип структуры
            Tags:     parseTagString(string(field.Tag)), // теги структуры
            Embedded: GetStructTags(field.Type), // рекурсивно вызываем для каждого поля эту же функцию; если поле — структура, то пройдёмся и по ней.
        }
    }

    return
}

// parseTagString десериализует тег-строку поля структуры.
// Дедупликация имён тегов: первый по порядку (слева направо).
// Ограничения: значение тега не может содержать символы ':' и '"'.
func parseTagString(tagRaw string) (retInfos TagsInfo) {
    retInfos = make(TagsInfo)

    // пример строки: json:"name" pg:"nullable,sortable"
    for _, tag := range strings.Split(tagRaw, " ") {
        if tag = strings.TrimSpace(tag); tag == "" {
            continue
        }

        tagParts := strings.Split(tag, ":")
        if len(tagParts) != 2 {
            continue
        }

        tagName := strings.TrimSpace(tagParts[0])
        if _, found := retInfos[tagName]; found {
            continue
        }

        tagValuesRaw, _ := strconv.Unquote(tagParts[1])
        tagValues := make([]string, 0)
        for _, value := range strings.Split(tagValuesRaw, ",") {
            if value := strings.TrimSpace(value); value != "" {
                tagValues = append(tagValues, value)
            }
        }

        retInfos[tagName] = tagValues
    }

    return
}
```

Example of execution:

```go
type (
    TestStruct struct {
        Id        string `json:"id" format:"uuid" example:"68b69bd2-8db6-4b7f-b7f0-7c78739046c6"`
        Name      string `json:"name" example:"Bob"`
        Group     Group  `json:"group"`
        CreatedAt int64  `json:"created_at" format:"unix" example:"1622647813"`
    }

    Group struct {
        Id             uint64   `json:"id"`
        PermsOverrides []string `json:"overrides" example:"USERS_RW,COMPANY_RWC"`
    }
)

func main() {
    var s *TestStruct
    fmt.Println(GetStructTags(s))
}
```

output:

```
{
   "CreatedAt": {
      "type": "int64",
      "tags": {
         "example": [
            "1622647813"
         ],
         "format": [
            "unix"
         ],
         "json": [
            "created_at"
         ]
      }
   },
   "Group": {
      "type": "main.Group",
      "tags": {
         "json": [
            "group"
         ]
      },
      "embedded": {
         "Id": {
            "type": "uint64",
            "tags": {
               "json": [
                  "id"
               ]
            }
         },
         "PermsOverrides": {
            "type": "[]string",
            "tags": {
               "example": [
                  "USERS_RW",
                  "COMPANY_RWC"
               ],
               "json": [
                  "overrides"
               ]
            }
         }
      }
   },
   "Id": {
      "type": "string",
      "tags": {
         "example": [
            "68b69bd2-8db6-4b7f-b7f0-7c78739046c6"
         ],
         "format": [
            "uuid"
         ],
         "json": [
            "id"
         ]
      }
   },
   "Name": {
      "type": "string",
      "tags": {
         "example": [
            "Bob"
         ],
         "json": [
            "name"
         ]
      }
   }
}
```

## Lesson Key Points

We've covered the basic uses of reflection in Go. Reflection allows you to determine the type and value of a variable based on its value. This allows you to create flexible code that works with a very wide range of input types. However, it's important to remember that reflection slows down your program. Whenever possible, you should use other approaches, as such a program will run extremely slowly.

We also explored:

- Determining the type of a variable.
- Determining the names and number of structure fields, and accessing them by index.
- Modifying a variable through reflection.
- Using tags.
