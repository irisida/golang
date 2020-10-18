package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("http interfaces example")

	resp, err := http.Get("https://google.co.uk")

	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}

	fmt.Println("response: ", resp.Body.Read)
}
