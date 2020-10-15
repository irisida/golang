![](https://hackernoon.com/drafts/0fnv29qd.png)

# Repo contents

- [Go - Developers guide](https://www.udemy.com/course/go-the-complete-developers-guide)
- [Master Go programming](https://www.udemy.com/course/master-go-programming-complete-golang-bootcamp).

# Go

## Go program structure

- package name
- imports
- package level variables and constants
- functions

```go
package main

import "fmt"

func main() {
    fmt.Println("This is a go program!")
}
```

## Packages

- all files in a package need to be in the same folder.
- each file in the package should declare its package name
- packages come in two types
  - executable (has a main)
  - library or reusable packages
- each project will have a package main because a main is required in an executable
  - within a main package will be a file with a `main()` function which is the automatic entry point of a program.

## Imports

imports are typically from:

- standard library
- custom imports, community packages or self defined packages
  - these type of packages are typically subject to a `go get` command and will require a relative path to them.

## Rules of assignment

- long form: `var varname type = value` or `var score float64 = 6.5`
- short form: `varname := value` or `score := 6.5`.
- inferred form: `var varname = value` or `var x = 6.5`
- All literals are untyped comstants, if numerically outside the binds of the intended type you will have an overflow complication error.

## Rules of constants

- You cannot change a constant.
- Constants belong to compile time, they cannot be initlaised at runtime.
  - One caveat being built-in functions such as len()
- You cannot use a variable to initialise a constant.

## The iota type

- Can be used a constant sequence/incrementor, or negated to be a decrementor.

## Data types

- `int8`, `int16`, `int32`, `int64` used to represent signed integers
- `uint8`, `uint6`, `uint32`, `uint64` for unsigned (positive) integers
- depending on platform `uint` is an alias for `uint32` or `uint64`
- depending on platform `int` is an alias for `int32` or `int64`
- `float32`, `float64` will drop a leading zero before the decimal point.
- `complex64` & `comple128`
- `byte` (alias for `uint8`)
- `rune` (alias for `int32`)
- `bool` predefined constants of `true` & `false`
- `string` type. unicode characters enclosed in double quotes. can be an emoty sequence.
- `array` & `slice` array is a numbered (indexed) sequence of same type elements. An array has a fixed length that is known at the time of declaration. A `slice` is dynamic allowing for shrinkage and growth.
- a `map` is an unordered group of elements of one type indexed by a set of keys of another type. (essentially it's the same as a dictionary in python)
- `struct` a sequence of named elements each of which has a name and a type, a class from OOP by stealth.
- `pointer` a pointer is a variable that stores the memory address of another variable. The value of an uninitialised pointer is nil.
- `function and interface`
- `channel` provides a mechanism for concurrently executing functions to communicate by sending and receiving values of a specified element type.
