package main

import "fmt"

func main() {

	xi := []int{2, 3, 4, 5}

	x := sum(xi...)

	fmt.Println(x)
}

func sum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}

	return sum
}
