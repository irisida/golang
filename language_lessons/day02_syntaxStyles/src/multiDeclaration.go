package main

import "fmt"

func multipleDeclaration() {

	/*
		here we can see the multiple variables listed, with the
		intended values following the same sequence.
	*/

	car, transmission, topSpeed, price := "720s", "Automatic", 210, 215999.00

	fmt.Printf("Car          - type is: %T \tvalue is: %v \n", car, car)
	fmt.Printf("transmission - type is: %T \tvalue is: %v \n", transmission, transmission)
	fmt.Printf("topSpeed     - type is: %T \tvalue is: %v \n", topSpeed, topSpeed)
	fmt.Printf("price        - type is: %T \tvlaue is: %v \n", price, price)

}
