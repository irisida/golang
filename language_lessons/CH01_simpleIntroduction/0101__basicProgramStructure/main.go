/*
	the very basic structure of a go program
	covering the:
		package
		imports
		main func
*/

package main

import "fmt"

/*
	set a package level demo for the usage of constant
	values. Like for all other programming languages
	a constant value once declared may not be changed
	programatically or by reassignment on the part of
	the programmer directly.
*/
const secondsInHour = 3600
const hoursInDay = 24

/*
	in the main package we are creating an executable
	package. It, thereore, expects there to be a main
	func that the compiler automatically understands
	as the execution entry point of any go program
*/
func main() {

	secondsInDay := secondsInHour * hoursInDay
	fmt.Printf("type: %T  value: %v", secondsInDay, secondsInDay)
}
