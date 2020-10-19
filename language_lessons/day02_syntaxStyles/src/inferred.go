package main

import "fmt"

func inferredVAriables() {

	// concise variables declaration
	x := 7
	y := 345.99
	z := "boom boom boom, let me hear you say whey-oh"

	fmt.Printf("x is: %v and of type: %T", x, x)
	fmt.Printf("y is: %v and of type: %T", y, y)
	fmt.Printf("z is: %v and of type: %T", z, z)
}
