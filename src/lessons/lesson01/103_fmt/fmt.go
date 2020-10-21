package main

import "fmt"

/*
	Package fmt:

	1) implements formatted I/O with functions analogous to the C language
	printf and scanf functions. It's used mainly to print to the stdout.
	2) fmt.Println() writes to standard output. Spaces are always added
	between operands and a newline is appended.
	3) fmt.Printf() prints out to stdout according to a format specifier
	called verb. It doesn't add a newline (\n)
		VERBS:
			%d -> decimal
			%f -> float
			%s -> string
			%q -> double-quoted string
			%v -> value (any)
			%#v -> a Go-syntax representation of the value
			%T -> value Type
			%t -> bool (true or false)
			%p -> pointer (address in base 16, with leading 0x)
			%c -> char (rune) represented by the corresponding Unicode code point
*/

func main() {
	var message string
	message = "Hello from Go"

	fmt.Println(message)

	x, y, z := 100, -100, .100

	fmt.Printf("x: %v  %T", x, x)
	fmt.Printf("y: %v  %T", y, y)
	fmt.Printf("z: %v  %T", z, z)
}
