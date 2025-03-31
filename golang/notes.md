# Go

## Introduction

Everything in Go is based of on packages it helps to modularize the code. `package main` contains the main function. `fmt` is a package for formatted I/O (printing).

### Hello World

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

Code is executed directly with `go run <program>.go`. However, we can build the program with `go build <program>.go` and run it with `./<program>`

### The Go-file-structure

There is three main folders: `src`, `bin` and `pkg`. The `src` folder contains the source code, the `bin` folder contains the compiled binaries for each project and the `pkg` folder contains the objects to be linked to create a self-contained binary.

Every time you compile source code, the Go linker looks inside the objects folder and links them together. Then ir creates the binary in the bin folder. This is one of the main advantages of Go, it's self contained.

```
/go
  /src
    /your-project-1
    /your-project-2
  /bin
    /compiled-binary-1
    /compiled-binary-2
  /pkg
    /objects-to-be-linked
```	

## Data Types and Control Structure

There is three types of data types in Go: atomic (typical data types), unsafe (pointers) and abstract (interfaces).

**Declaring a variable**
Simply declare the variable with `var <name> <type>` (e.g. `var x int`). To declare multiple variables, use `,` (e.g. `var x, y int`) or parenthesis (e.g. `var (x = 1, y = 2)`). When value is assigned in the declaration, the type is inferred from the value (e.g. `var x = 1`), it isn't necessary to specify the type.
To cast a value to a specific type, use `(<type>) <value>` (e.g. `(int) "1"`). 
Finally, rather than declaring variables, you can use `:=` to declare and assign a value to a variable (e.g. `x := 1`).

### Strings

All string in Go are mutable, it means that you can change their value string again and again. To define a string just `var <name> string` and to assign value, use `= "<value>"` (e.g. `x = "Hello"`). 

The Go `:=` operator can be only used for declaring variables, not for assigning values to them. So, if we define a string using `:=`, we cannot change it's value again using `:=`, we can use `=` instead.

**Strings functions**
- `len(<string>)` returns the length of the string
- `<string>[<index>]` returns the character at the specified index
- `.contains(<string>, <string>)` returns true if the second string is contained in the first string
- `.replace(<string>, <string>, <string>)` returns the string with the first string replaced by the second string
- `.split(<string>, <string>)` returns an array of strings separated by the second string
- ...

### Arrays

List of elements of the same type. To declare an array, use `var <name> [size] <type>` (e.g. `var x [3]int`). To declare and assign a value to an array, use `:=` (e.g. `x := [3]int{1, 2, 3}`).

**Arrays functions**
- `len(<array>)` returns the length of the array
- `append(<array>, <value>, <value>, ...)` returns a new array with the value appended
- ...

### Pointers

A pointer is a variable that holds the memory address of another variable. To declare a pointer, use `var <name> *<type>` (e.g. `var x *int`) and then assign the memory address of another variable using `&<variable>` (e.g. `x = &y`). In this example `x` is a pointer to `y`.

We can use pointers to access the value of another variable. To do so, use `*<pointer>` (e.g. `*x`). They are useful when using functions that modify the value of another variable.

### Structs

A struct is a collection of fields. To declare a struct, use `type <name> struct { <field> <type> }` (e.g. `type Person struct { name string, age int }`). To declare and assign a value to a struct, use `:=` (e.g. `p := Person{name: "John", age: 30}`).

**Struct functions**
- `<struct>.<field>` returns the value of the field
- `<struct>.<field> = <value>` sets the value of the field
- ...

It can be defined functions on structs. For example: 
```go
func (p Person) sayHello() { 
    fmt.Println("Hello, my name is", p.name) 
}
```