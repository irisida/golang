package main

import "fmt"

/*
	coverage:
	1. 	long form variable declaration - this is the more formal type
		of declaration and we'll see the var keyword, the data type and
		some initialisation. It's worth noting that when package level
		variables are declared you cannot use the short form we will
		see in this file. the must have the longer form like the
		example below.

	2.	short form - here we see a more
	demonstrating a multifile program where we have kept the
	files as part of the main package. multipackage programs
	will come later :)

*/

func main() {
	demoVariableDeclarations()
	multipleDeclaration()
}

func demoVariableDeclarations() {
	// long form variable declaration, note we see
	// the use of var, the identifier, the datatype
	// and the initialisation value.
	var x int = 7

	// non initialised long form.
	// declaration is complete and the initialisation
	// is done on the following line
	var s1 string
	s1 = "Go programming"

	// set the output
	fmt.Println("long form  : x = ", x)
	fmt.Println("long form  : s1 = ", s1)

	// The short declaration method. Note that the short
	// assignment method requires us to use the colon ":".
	// Only on the first assignment operation to a variable
	// and any reassignment can omit the colon.
	age := 30

	fmt.Println("short form : age = ", age)
	fmt.Printf("short form : type of age = %T \n", age)
	fmt.Println("short form : x = ", x)
}
