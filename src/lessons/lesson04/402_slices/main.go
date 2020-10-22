package main

import "fmt"

func main() {

	// declaring a slice
	x := []int{1, 3, 5, 7, 9}
	fmt.Printf("%#v \n", x)

	// appending to a slice
	x = append(x, 11)
	fmt.Printf("%#v \n", x)

	// slicing a slice
	// remember that slices are zero indexed. Slicing a slice can
	// have multiple parts, sliceName[startposition : stopPosition]
	// start position is inclusive, if omitted it is assumed zero
	// stop position is exclusive, if omitted it is assumed as the
	// last element of the slice.
	y := x[2:4]
	fmt.Printf("%#v \n", y)

	// deleting from a slice
	// there is no remove/delete function so we can co=opt append to
	// achieve a delete operation.
	x = append(x[:1], x[len(x)-1:]...)
	fmt.Printf("%#v", x)
}
