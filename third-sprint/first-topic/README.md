# Summary

OOP mechanisms in Go:

- **Abstraction**. Structures with fields and methods allow you to describe an object as a logically complete unit. This mechanism makes it possible to build more complex abstractions from component parts.
- **Encapsulation**. Structure methods, with the ability to specify their visibility (public/private), allow you to hide the implementation under the hood, simplifying user interaction with the object. Go does not have the concept of an "object constructor," and this can cause some inconvenience for the user of a type. If an object type is exported, the user can create it manually, bypassing the written "constructor" NewMyType(), and get an object that is correct from the compiler's perspective, but invalid from the application's logic perspective.
- **Polymorphism**. The concept of polymorphism in Go is implemented using interfaces. They allow you to define a type specification without binding it to its implementation. Interfaces will be discussed in the next topic.
- **Inheritance**. Nested structures don't allow for a full implementation of inheritance. Standard methods (without reflection) can't check whether two objects have a "child-parent" relationship. Go doesn't have an equivalent of an "abstract class," so it's impossible to take a pointer to a child's base type (for example, if a function requires a base type as input). These limitations can be overcome by using interfaces, which allow you to check whether an object implements a specific specification.

Go has its own vision of standard OOP mechanisms. These implementation peculiarities don't limit Go developers in their coding; rather, they require a different perspective on OOP standards.
