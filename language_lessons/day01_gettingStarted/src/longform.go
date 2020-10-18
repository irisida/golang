package main

import "fmt"

func longFormVariables() {
	/* method 1 using the var keyword */
	var x int = 7

	/* non initialised method 1  */
	var s1 string
	s1 = "Go programming"

	fmt.Println("method 1: x = ", x)
	fmt.Println("method 1: s1 = ", s1)
}
