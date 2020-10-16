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
