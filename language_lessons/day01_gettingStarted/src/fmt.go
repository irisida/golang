package main

/*
	Package fmt implements formatted I/O with functions
	analogous to C's printf and scanf. It's used mainly
	to print out to stdout
*/

import "fmt"

func fmtExample() {

	/*
		fmt.Println() writes to standard output.
		spaces are always added between operands and a newline is appended.
	*/
	fmt.Println("Typical hello message from some programming language") // => Hello Go World!

	var name, age = "Sherlock", 40
	fmt.Println(name, "is", age, "years old.")

	/*
		fmt.Printf() prints out to stdout according to a format specifier called verb.
		It doesn't add a newline (\n)

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
	a, b, c := 10, 15.5, "Gophers"
	grades := []int{10, 20, 30}

	fmt.Printf("a is %d, b is %f, c is %s \n", a, b, c)
	fmt.Printf("%q\n", c)
	fmt.Printf("%v\n", grades)
	fmt.Printf("%#v\n", grades)
	fmt.Printf("b is of type %T and grades is of type %T\n", b, grades)

	const pi float64 = 3.14159
	fmt.Printf("pi is %.4f\n", pi) // formatting with 4 decimal points

	fmt.Printf("255 in base 2 is %b\n", 255)  //  => 255 in base 2 is 11111111
	fmt.Printf("101 in base 16 is %x\n", 101) // => 101 in base 16 is 65

}
