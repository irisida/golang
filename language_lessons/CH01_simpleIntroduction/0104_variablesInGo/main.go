package main

import "fmt"

func shortFormVariables() {
	/*
		method 2 the short declaration method. Note that the
		short assignment method requires us to use the colon
		only on the first assignment operation to a variable
		and any reassignment can omit the colon.
	*/
	age := 30

	fmt.Println("method 2: age = ", age)
	fmt.Printf("method 2: type of age = %T \n", age)
}

func longFormVariables() {
	/* method 1 using the var keyword */
	var x int = 7

	/* non initialised method 1  */
	var s1 string
	s1 = "Go programming"

	fmt.Println("method 1: x = ", x)
	fmt.Println("method 1: s1 = ", s1)
}

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

func main() {
	longFormVariables()
	shortFormVariables()
	multipleDeclaration()
}
