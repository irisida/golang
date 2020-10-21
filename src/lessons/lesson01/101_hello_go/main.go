package main

import "fmt"

/*
	The fmt package offers several functions that take an empty
	interface of variadic inputs. This means we can submit and
	expect to see the output of mixed types.
*/

func main() {

	// classic hello world example
	fmt.Println("hello, World")

	// demonstration of the empty interface taking multiple types
	fmt.Println(0.28, 1000, "Hi from Go")

	// same interface called with different set of values
	// this is unlimited.
	fmt.Println(10, 9, 8, 7.5, "Six, five four", 3, 2, "one")
}
