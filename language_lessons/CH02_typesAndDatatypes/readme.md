![](https://hackernoon.com/drafts/0fnv29qd.png)

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
