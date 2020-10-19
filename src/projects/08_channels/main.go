package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://google.co.uk",
		"https://facebook.com",
		"http://stackoverflow.com",
		"https://amazon.co.uk",
		"http://golang.org",
	}

	for _, link := range links {
		//go checkLink(link)
		checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("❌ : %s could be down \n", link)
		return
	}

	fmt.Printf("✅ : %s \n", link)
}
