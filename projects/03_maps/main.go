package main

import "fmt"

func main() {
	/* maps literal syntax */
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	/* using make() */
	makeMap := make(map[string]string)

	makeMap["Ireland"] = "Dublin"

	fmt.Println(colors)
	fmt.Println(makeMap)

	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
