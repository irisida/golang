package main

import "fmt"

func main() {

	/*
		multi dimensional demo.
		here we have multiple dimensions of slice of string
	*/

	rec1 := []string{"userid", "https://google.co.uk", "http://stackoverflow.com"}
	rec2 := []string{"userid", "Monday", "Wednesday", "Friday"}

	recs := [][]string{rec1, rec2}
	fmt.Println(recs)
}
