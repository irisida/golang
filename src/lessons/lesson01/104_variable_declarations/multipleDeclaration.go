package main

import "fmt"

func multipleDeclaration() {

	// concise way to declare multiple variables with inferred types.
	// type inference is where the Go compiler is allowed to derive the
	// datatype of a value.
	car, transmission, topSpeed, cost := "720s", "Automatic", 210, 215999.00

	fmt.Printf("Car: value: %v type: %T \n", car, car)
	fmt.Printf("transmission: value: %v type: %T \n", transmission, transmission)
	fmt.Printf("topSpeed: value: %v type: %T \n", topSpeed, topSpeed)
	fmt.Printf("price: value: %v type: %T \n", cost, cost)

	// another way to demonstrate collection of variables being declared
	// note the single use of var, also note the implied connection between
	// variables. This is deceptive as no relationship is implied, this is
	// not a struct (class) and each item has no relation to its neighbour.
	var (
		screen    string
		processor string
		memory    int
		price     float64
	)

	// lets see the zero values. A zero value is where a variable that
	// has not been assigned a value is used with its default value that
	// is assigned by Go.
	fmt.Printf("%#v %#v %#v %#v", screen, processor, memory, price)

}
