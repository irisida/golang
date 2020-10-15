/*
	the very basic structure of a go program
	covering the:
	package
	imports
	main func
*/

package main

import "fmt"

/* set a package level constat value */
const secondsInHour = 3600

/* main execution entry point of any go program */
func main() {
	fmt.Println("Hello Go")
	distance := 60.8

	num := 107

	fmt.Printf("Distance in miles is %f \n", distance*0.62137)
	fmt.Printf("%x", num)
}
