package main

import (
	"fmt"

	"github.com/irisida/golang/src/lessons/examples/named_types_part3/weights"
)

func main() {

	type (

		// Gram underlying type is int
		Gram int

		// Kilogram underlying type is int
		Kilogram Gram

		// Ton underlying type is int
		Ton Kilogram
	)

	var (
		salt   Gram     = 100
		apples Kilogram = 5
		truck  Ton      = 10
	)

	fmt.Printf("salt: %d, %T \n", salt, salt)
	fmt.Printf("apples: %d, %T \n", apples, apples)
	fmt.Printf("truck: %d, %T \n", truck, truck)

	salt = Gram(weights.Gram(1000))
	fmt.Printf("salt: %d, %T \n", salt, salt)

}
