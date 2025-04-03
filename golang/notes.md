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

### Interfaces

An interface is a set of methods. It's useful to define a common behavior for a set of types. To declare an interface, use:

```go
type <name> interface {
    <method>(<parameters>) <return type>
    <method>(<parameters>) <return type>
    ...
}
```

Inside the file we can define data types and the behavior of the functions defined in the interface. We can define the same function with different parameters and return types.

One method to ensure that all the structures implement the interface methods is using a code like this:

```go
func New<name>(arg) <name> {
  return &<structName>{arg}
}
```

If we create a new interface object and return it as a pointer to the struct, Go will check if the struct implements the interface methods and if not, it will panic.

**Empty interface**

An empty interface is a type that can hold any value. It's a black box. This is very helpful when we are talking about development because there most of the time you don't know what type of data you are going to receive.

```go
interface {}

func Anything(anything interface{}) {
  fmt.Println(anything)
}
```

For example we can use this for maps, telling tha the key is a string and the value can be any type.

```go
mymap := make(map[string]interface{})

mymap["name"] = "John"
mymap["age"] = 30

fmt.Println(mymap)
```

### Control Structures

The main control structures in Go are:

- `if <condition> { <code> }`
- `if <condition> { <code> } else { <code> }`
- `if <condition> { <code> } else if <condition> { <code> } else { <code> }`
- `switch <expression> { case <value>: <code> ... default: <code> }`
- `for <condition> { <code> }`
- `while <condition> { <code> }`
- `do { <code> } while <condition>`
- `break`
- `continue`

Note Go doesn't use parentheses in the conditions.

`switch` can also be defined without the expression, e.g. `switch { case <expression>: <code> ... default: <code> }`. In this case, the expression is the condition defined in the case statement.

## Concurrency

Concurrency is the ability to execute multiple tasks or processes at the same time. It is a way to improve the performance of a program by using multiple cores or threads.

### Introduction

Golang processes don't run on threads but on goroutines. A goroutine is a light-weight thread of execution.

### Synchronous vs Asynchronous

The main difference between synchronous and asynchronous programming is that synchronous programming is executed in a sequential order, while asynchronous programming is executed in parallel.

### Goroutines

A goroutine is a light-weight thread of execution. It is a function that can be executed concurrently with other goroutines.

To create a goroutine, use `go <function>()`. For example: `go sayHello("John")`. This will run the function in a separate goroutine, in the background. For example if we run a goroutine with a sleep, the main goroutine will continue to run while the goroutine is sleeping. Or if we have a `print` in the gorutine after the `sleep`, it won't be printed through the console because the main goroutine has already finished.

**Waitgroups**

To avoid this we can use `waitgroups` to wait for the goroutine to finish. For example: `wg.Wait()`. This will wait for all the goroutines to finish before continuing the main goroutine.

To define a waitgroup, we must declare three functions: `add`, `done` and `wait`. The `add` function is used to add a goroutine to the waitgroup, the `done` function is used to notify that the goroutine has finished, and the `wait` function is used to wait for all the goroutines to finish.

```go	
wg := &sync.WaitGroup{}

wg.Add(1)   // Add 1 to the WaitGroup
go func() {
  fmt.Println("Hello, World!")
  wg.Done() // Notify that the goroutine has finished
}
wg.Wait()   // Wait for all the goroutines to finish
```

Waits can be used to control the synchronization of goroutines. However, if we want to control the manipulation of a shared resource, we can use mutexes.

**Mutexes**

A mutex is a lock that is used to control the access to a shared resource. It is useful when we want to control the access to a shared resource. For example:

```go
mutex := &sync.Mutex{}
mutex.Lock()
go func() {
  // Do something with the shared resource
}()
mutex.Unlock()
```

**Lambda Functions**

Lambda functions are functions without a name. They are useful when we want to create a function that is only used once. For example: 

```go
func() { 
  fmt.Println("Hello, World!") 
  }()
```
