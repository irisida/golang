package main

import "fmt"

func main() {

	/* 	a defined type or named type is basically a
	 	label over a primitive type. This underlying
		is the real type, but we have applied a meaninful
		name to the type which increases comprehension
		of complex or abstract ideas.
	*/

	// example
	type WeeNumber uint8

	var counter WeeNumber

	fmt.Printf("%T %v \n", counter, counter)

	counter = 255

	fmt.Printf("%T %v \n", counter, counter)

	counter++

	fmt.Printf("%T %v \n", counter, counter)

	// we can see the same uint8 rules

	// lets see a small caveat here.
	var otherCounter uint8 // same type as WeeNumber's underlying type
	otherCounter = 10

	// note that even though they share the same underlying type
	// they siull cannot be used completely interchangably.
	// counter = otherCounter (type compilation error is thrown)
	counter = WeeNumber(otherCounter) // applies a conversion
}
