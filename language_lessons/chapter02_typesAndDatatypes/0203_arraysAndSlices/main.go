package main

import "fmt"

func main() {

	// array type
	var nums = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Printf("%T \n", nums)
	fmt.Println(nums)

	// aray type of string
	var cities = [5]string{"Glasgow", "Dublin", "London", "Luxembourg", "Southport"}

	fmt.Printf("%T \n", cities)
	fmt.Println(cities)

	/*
		slice type of string, note the number of elements
		is not prescribed and this allows for dynamicism
	*/
	var moreCities = []string{"Glasgow", "Dublin", "London", "Luxembourg", "Southport"}

	fmt.Printf("%T \n", moreCities)
	fmt.Println(moreCities)

	// map type
	balances := map[string]float64{
		"USD": 0.165,
		"GBP": -0.85,
		"Eur": 1.113,
		"AUD": -2.03,
	}

	fmt.Printf("%T \n", balances)
	fmt.Println(balances)
}
