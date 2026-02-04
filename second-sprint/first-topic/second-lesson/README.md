# Second lesson

## Arrays

Syntax:

```go
// declaration
var <name> [<length>]<type>

// initialisation
<name> := [length]<type> {array elements}
```

In Go, the number of elements in an array is a part of the type, i.e. arrays `[3]int` and `[5]int` belong to different types.

The compiler monitors the size of the initialization list, and it will not be possible to compile such a construction:

```
rgbColor := [3]uint8{1, 2, 3, 4} // array index 3 out of bounds [0:3]
```

The number of elements in the array can be output automatically based on the length of the initialization list. The following construction is used for this:

```go
rgbColor := [...]uint8{255, 255, 128} // [255 255 128] len = 3
rgbaColor := [...]uint8{255, 255, 128, 1} // [255 255 128 1] len = 4
```

Three dots indicate to the compiler that the size of the array should be derived based on the size of the initialization list.
Sometimes it becomes necessary to specify only one or more array elements in the initialization list, and leave the others untouched.
If you needed to specify the average daily temperature on Sunday, the code would hardly be pleased with its sophistication:

```go
thisWeekTemp := [7]int {0,0,0,0,0,0,11} // [0 0 0 0 0 0 11]
```

For large arrays, this would turn into an unreadable construct:

```go
var thisWeekTemp [7]int // [0 0 0 0 0 0 0]
thisWeekTemp[6] = 11 // [0 0 0 0 0 0 11]
```

However, only the necessary elements and their indexes can be specified in the initialization list. The index and value are separated by a colon:

```go
thisWeekTemp := [7]int {6:11, 2:3} // [0 0 3 0 0 0 11]
```

The size of the array can be obtained by the built-in `len` function. Since the size of the array is known at the compilation stage, the calculation of this function is replaced by a specific value during compilation.

If you try to access an array element beyond its size, the memory protection mechanism in Go will trigger and panic will begin. Unlike the C language, the Go compiler and runtime control going beyond the array, preventing access to invalid memory areas.

## Multidimensional arrays

Syntax:

```go
var <name> [<rows>][<columns>]<type>

var rgbImage [1080][1920][3]uint8
line := rgbImage[2]
pixel := rgbImage[2][3]
colour := rgbImage[2][3][1]
```

## Traversing array values

Example:

```go
var weekTemp = [7]int{5, 4, 6, 8, 11, 9, 5}

sumTemp := 0

for i:= 0; i < len(weekTemp); i++ {
    sumTemp += weekTemp[i]
}

average := sumTemp / len(weekTemp)
```
Go has a more convenient `for range` (the idiomatic equivalent of `for-each`) construct that allows you to traverse array elements sequentially without using additional variables:

```go
var weekTemp = [7]int{5, 4, 6, 8, 11, 9, 5}

sumTemp := 0

for _, temp := range weekTemp {
    sumTemp += temp
}

average := sumTemp / len(weekTemp)
```

The `range` operator returns the index and value of the next element in the array at each iteration.

**Note the following important point: the `_, temp := range weekTemp` construction creates a new `temp` variable, the type of which will be determined by the type of the array element.**

This variable will be assigned the next value from the array at each iteration of the loop. If you change the value of the `temp` variable, it will not affect the values in the array.

To access an array element, you will need an index:

```go
var weekTemp = [7]int{5, 4, 6, 8, 11, 9, 5}
for _, temp := range weekTemp {
   temp = 0
}
// weekTemp [5 4 6 8 11 9 5]
// if the value of the element is not used, you can omit the second variable
for i := range weekTemp {
   weekTemp[i] = 0
}
// weekTemp [0 0 0 0 0 0 0]
```

Array variables can be assigned to each other, but they must have the same type, and the number and type of elements must match.
During the assignment process, a complete copy of the array is performed, and if the program processes sufficiently large arrays of data, these copies can significantly slow down the program and increase memory consumption.
**Passing an array** (as well as a variable of any other type) to a function is copying its value to a variable of the function argument. There will also be a full copy here.

**There is another important point associated with the `for-range` loop — the `range` operand is copied to a temporary variable that is already being used for traversal.**

This can also slow down the execution of the program for arrays, so you should use pointer capture to avoid being idle:

```go
for i, temp := range &weekTemp {
    fmt.Println(i, temp)
}
```

### Advantages of using arrays

    - The array elements are always sequentially located in memory, which makes the processor happy and speeds up program execution.
    - Arrays have a fixed length, so the allocation of memory for an array occurs exactly once at the time of its declaration.
    - The access time to the array elements is minimal.
    - Go checks for going outside the array at the compilation stage if it can calculate the index value of an element at the compilation stage and during program execution. In the first case, there will be a compilation error, and in the second, there will be panic. It is better not to allow panic.

### Disadvantages of using arrays

    - Arrays can only be of a fixed length: if the number of elements is unknown to us in advance, we will have to allocate memory with a margin.
    - Arrays are transferred and assigned with a complete copy of the elements, which can lead to sudden performance degradation and increased memory consumption.
    - To process arrays of different sizes, you will have to write different functions (if generics are not used).

Arrays should be used very deliberately, when the size of your array is precisely known at the compilation stage. This allows you to speed up the program.
To fix the disadvantages of arrays, Go introduced slices.

## Slices

Since the length of an array is part of its type, arrays are not suitable for storing dynamic—sized data collections. This problem is solved by another type of data, which is much more often used in practice, slice.

In addition, the slice will save us from the problem of copying arrays for assignment.

A slice is a sequence of variable length consisting of elements of the same type. The slice type is written as an array type without specifying the size. You can initialize a "slice" type variable with values, but unlike an array, a variable without initialization is `nil`.

```go
var mySlice []int
```
A slice is very similar to a `list` in Python, but it has its own characteristics.
A slice is a wrapper over an array pointer, and in Go a slice is used as a structure of the following type:

    - the pointer to the first element of the base array is ptr;
    - slice length — len, the number of elements in the slice;
    - slice capacity — cap, the number of elements in the array.

![assets](assets/1.png)

The slice parameters `len` and `cap` can be obtained by calling the corresponding built-in functions `len()` and `cap()`.
If you simply declare such a structure, it will be equal to `nil`.

The built-in `make()` function is used to create a slice:

```go
    mySlice := make([]TypeOfelement, LenOfslice, CapOfSlice)
    mySlice := make([]int, 0)
    mySlice := make([]int, 5)
    mySlice := make([]int, 5, 10)
```

Arguments to the `make` function:

    - The type of slice (empty square brackets and the type of the slice element).
    - Slice length. If it is not passed, it is equal to zero by default.
    - The slice capacity is the size of the base array. If no value is passed, it defaults to the length of the slice.

The `make` function creates an array with a length of `cap` and writes a pointer to it into the slice structure. It also fills in the `len` and `cap` fields in this structure and returns it as a "slice" type variable.
Even if `len` and `cap` are passed as null, the structure itself will no longer be equal to `nil`. It will be allocated in memory, and the pointer to the base array will get a value other than `nil`.
If you pass the `cap` parameter less than `len` to the make function, a compilation error or a panic during execution will be caused.

A slice can be created from a composite literal in the same way as an array. The only difference is that we don't specify the size of the array:

```go
s := []int{1, 2, 3}  // [1 2 3]
```

The length and capacity of the slice will be equal to the composite literal.

A new slice can be created based on an existing slice or array. For this, the slice capture operation is used.
It is performed using two brackets with a colon `[i:j]`, where `i` is the index of the first element of the new slice, and `j` is the index of the next element **not included** in the new slice.

It is allowed not to specify `i` and `j`. In this case, `i` will be 0 by default, and `j` will be equal to the length of the array or slice.

Thus, `[:]` will return a slice of the entire array, `[:k]` — from the beginning to the kth element, `[k:]` - from the kth element to the end of the array.

`i` and `j` must be non-negative and no greater than `len`, with `i` being less than or equal to `j`. If these conditions are not met, a compilation error or panic will occur.

Consider an example:

```go
    weekTempArr := [7]int{1, 2, 3, 4, 5, 6, 7}
    workDaysSlice := weekTempArr[:5]
    weekendSlice := weekTempArr[5:]
    fromTuesdayToThursDaySlice := weekTempArr[1:4]
    weekTempSlice := weekTempArr[:]

    fmt.Println(workDaysSlice, len(workDaysSlice), cap(workDaysSlice)) // [1 2 3 4 5] 5 7
    fmt.Println(weekendSlice, len(weekendSlice), cap(weekendSlice)) // [6 7] 2 2
    fmt.Println(fromTuesdayToThursDaySlice, len(fromTuesdayToThursDaySlice), cap(fromTuesdayToThursDaySlice)) // [2 3 4] 3 6
    fmt.Println(weekTempSlice, len(weekTempSlice), cap(weekTempSlice)) // [1 2 3 4 5 6 7] 7 7
```

## Changing the slice size

A slice is a sequence of dynamically sized elements.
The slice size is reduced through the operation of taking a slice. The result of the capture can be assigned to the same slice:

```go
    s := []int{1,2,3} // [1 2 3]
    s = s[:len(s)-1] // [1 2]
```

The capacity of the array does not change.
The built-in append function is used to add elements to a slice. It accepts a variable of the "slice" type and one or more variables of the slice element type, then returns a new slice, which consists of a copy of the passed slice and the elements passed into it.

**Attention: append does not modify the slice passed to it, but creates a new one based on the passed one.**

Consider example:

```go
    a := []int{1, 2, 3, 4}
    b := a[2:3]   // b = [3]
    b = append(b, 7)
    fmt.Println(a, len(a), cap(a)) // [1 2 3 7] 4 4
    fmt.Println(b, len(b), cap(b)) // [3 7] 2 2
    b = append(b, 8, 9, 10)
    b[0] = 11
    fmt.Println(a, len(a), cap(a)) // [1 2 3 7] 4 4
    fmt.Println(b, len(b), cap(b)) // [11 7 8 9 10] 5 6
```

Slice `a` changed after executing the first `append`, because slice `b` increased its length, but did not go beyond the base array.
The point is in the mechanics of `append`: if the slice capacity allows you to place the added elements (that is, the difference between the length of the slice and its capacity is greater than or equal to the number of placed elements), then `append` returns a new slice, which is obtained from an existing slice by adding new elements inside the base array.
If the slice's capacity does not allow these elements to be placed, then a new base array of suitable size is created, all the elements of the transferred slice are copied into it and new ones are added. That is why in the example, after the second `append`, slice `b` refers to the new base array. The difference between the capacity of the new base array and its length depends on the number of elements. In the example, the length of the base array for `b` is 5, and the capacity is 6.

For 1000 items, the difference will be different:

```go
    a := make([]int, 1000)
    b := append(a, 7)
    fmt.Println(len(a), cap(a))  // 1000 1000
    fmt.Println(len(b), cap(b))  // 1001 1536
```

To connect two slices, you need to unpack the `append(a,b...)` slice. The function takes a number of individual elements and converts the slice into a list through unpacking.

Consider a few more examples:

```go
s := make([]int, 4, 7) // [0 0 0 0], len = 4 cap = 7
// 1. Create a slice s with a base array of 7 elements.
// The first four elements will be available in the slice.

slice1 := append(s[:2], 2, 3, 4)
fmt.Println(s, slice1) // [0 0 2 3] [0 0 2 3 4]
// 2. We take a slice from the first two elements of s and add three elements to them.
// Since the total length of the resulting slice (len == 5) is less than the capacity of s[:2] (cap == 7),
// then the base array remains the same.
// Slice s has also changed, but its length remains the same

slice2 := append(s[1:2], 7)
fmt.Println(s, slice1, slice2) // [0 0 7 3] [0 0 7 3 4] [0 7]
// 3. Here, too, the base array remains the same, all three slices have changed.

slice3 := append(s, slice1[1:]...)
fmt.Println(len(slice3), cap(slice3)) // 8 14
// 4. The length of s and slice1[1:] is 4, the length of the new slice will be 8,
// which is greater than the capacity of the base array.
// A new base array with a capacity of 14 will be created,
// the capacity of the new base array is selected automatically
// and depends on the current size and the number of added elements

// 5. It is easy to verify that slice3 refers to the new base array
s[1] = 99
fmt.Println(s, slice1, slice2, slice3)
// [0 99 7 3] [0 99 7 3 4] [99 7] [0 0 7 3 0 7 3 4]
```

This is how this process looks in the picture:

![assets](assets/2.png)

It's not very clear here whether the new slices will reference the same base array or send their copies to the new array. Therefore, in practice, the `append` function is recommended only for assigning a slice to itself: `s = append(s, b)`.

**To use slices, you need to understand how they work. Otherwise, you will get errors that are very difficult to find. Try to avoid situations where multiple slices link to the base array and elements are added or changed.**

The slice capture operation also supports the third parameter: `[low:high:max]` — the third parameter specifies the capacity of the base array required to create a new slice. In this case, the `max` must be less than or equal to the capacity of the base array or slice.

## Assigning a slice and passing it to a function

Assigning slice variables to each other, even the most impressive size, does not consume much computing power, because the slice structure itself always contains only three fields: `ptr`, `len`, and `cap`. However, we must keep in mind that these variables refer to the same array, so a change in the data of one slice may lead to a change in the other.

When passing a slice to the arguments of a function, the slice structure is copied to a local variable inside the function. This allows you to change the data inside the slice passed to the function. However, if you need to add or remove elements from a slice, then these changes will affect only the local variable of the slice.

For example, let's take a couple of functions from the standard library:

```go
    s := []int{5, 4, 1, 3, 2}
    sort.Ints(s)
    fmt.Println(s) // [1 2 3 4 5]
```

The sort function `sort.Ints()` the resulting slice of integers in ascending order. She does not change the size and capacity of the slice, so she can safely work with it.

```go
    import (
    "bytes"
    "fmt"
)

func main() {
    bSlice := []byte(" \t\n a lone gopher \n\t\r\n")
    fmt.Printf("%s", bytes.TrimSpace(bSlice)) // a lone gopher
    fmt.Printf("%s", bSlice)  // \t\n a lone gopher \n\t\r\n

}
```

The `bytes.TrimSpace` function takes a slice of bytes and returns a new slice of bytes from where the beginning and ending whitespace characters have been removed. The slice size should change, which means that the `bSlice` will remain intact. As a result, `bytes.TrimSpace` will give us a new slice.

## Copying slices

To copy elements from one slice to another, the `copy([]T dest, []T src)` function is used, where `dest` is the receiver slice and `src` is the source slice. This function only overwrites the elements, so the number of copied elements will be equal to the shorter length of the two slices.

```go
var dest []int
dest2, dest3 := make([]int, 3),  make([]int, 5)
src := []int{1, 2, 3, 4}
copy(dest, src)
copy(dest2, src)
copy(dest3, src)
fmt.Println(dest, dest2, dest3, src ) // [] [1 2 3] [1 2 3 4 0] [1 2 3 4]
```

### Slice traversal and access to elements

Slice traversal and access to slice elements occur in exactly the same way as for arrays. To get to the elements by index, square brackets `[]`, `for` and `for-range` loops are used.

### Useful techniques for working with slices

Unlike other programming languages, Go does not flaunt an abundance of functions for working with slices. The gophers got out and came up with several techniques that allow them to solve common problems.

Deleting the last slice element:

```go
    s := []int{1, 2, 3}
    if len(s) != 0 { // avoid panic
        s = s[:len(s)-1]
    }

    fmt.Println(s) // [1 2]
```

Deleting the first slice element:

```go
    s := []int{1,2,3}
    if len(s) != 0 { // avoid panic
        s = s[1:]
    }
    fmt.Println(s) // [2 3]
```

Deleting a slice element with index `i`:

```go
    s := []int{1,2,3,4,5}
    i := 2

    if len(s) != 0 && i < len(s) { // avoid panic
        s = append(s[:i], s[i+1:]...)
    }
    fmt.Println(s) // [1 2 4 5]
```

Comparing two slices:

```go

    s1 := []int{1,2,3}
    s2 := []int{1,2,4}
    s3 := []string{"1","2","3"}
    s4 := []int{1,2,3}

    fmt.Println(reflect.DeepEqual(s1,s2)) // false
    fmt.Println(reflect.DeepEqual(s1,s3)) // false
    fmt.Println(reflect.DeepEqual(s1,s4)) // true
```

## Conclusion

Slices are widely used in Go to implement collections of the same type of elements, but they require proper use: without understanding their structure, you can run into tricky mistakes.



#### P.S

How can I create a slice based on an array?

```go
slice := array[:]
```
