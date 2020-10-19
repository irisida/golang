package main

import (
	"fmt"
	"runtime"
)

func main() {

	/*
		1. runtime.NumCPU() is an expression, it returns a value.
		2. fmt.Printlin() is a statement. It tells Go to perform an action, to do something.
	*/

	fmt.Println(runtime.NumCPU())
}
