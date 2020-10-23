package main

import "fmt"

type gram float64
type ounce float64

func main() {

	var g gram = 1000
	var o ounce

	/*
		note that with named types, even when a named type
		has the same underlying type as another named type
		it cannot simply be considered as the same and must
		be converted to take advnatage of the operations.
		conversion is easy because the underlying types will
		match but the conversion must be done.

		below if we have:
		o = g * 0.035274 - it will fail, we muct convert it
		to the ounce type.
	*/

	o = ounce(g) * 0.035274

	fmt.Printf("%g grams is %.2f  ounces \n", g, o)
}
