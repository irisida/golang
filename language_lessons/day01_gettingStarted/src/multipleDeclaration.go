package main

import "fmt"

func multipleDeclaration() {
	car, transmission, topSpeed, price := "720s", "Automatic", 210, 215999.00

	fmt.Println("Car type is: ", car)
	fmt.Printf("Car variable type is: %T \n", car)

	fmt.Println("transmission type is: ", transmission)
	fmt.Printf("transmission variable type is: %T \n", transmission)

	fmt.Println("top speed is: ", topSpeed)
	fmt.Printf("top speed variable type is: %T \n", topSpeed)

	fmt.Println("price is: ", price)
	fmt.Printf("price variable type is: %T \n", price)
}
