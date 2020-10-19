package main

import "fmt"

func multipleDeclaration() {

	car, transmission, topSpeed, price := "720s", "Automatic", 210, 215999.00

	fmt.Printf("Car: value: %v type: %T \n", car, car)
	fmt.Printf("transmission: value: %v type: %T \n", transmission, transmission)
	fmt.Printf("topSpeed: value: %v type: %T \n", topSpeed, topSpeed)
	fmt.Printf("price: value: %v type: %T \n", price, price)
}
