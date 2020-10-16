package main

import (
	"fmt"
	"runtime"
)

func main() {

	/*
		runtime.NumCPU() is an expression, it returns a value.
		fmt.Printlin() is a statement. It tells Go to perform
		an action, to do something.
	*/

	fmt.Println(runtime.NumCPU())
}
