package main

/*
	go routine and channels example.
*/

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.co.uk",
		"https://facebook.com",
		"http://stackoverflow.com",
		"https://amazon.co.uk",
		"http://golang.org",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		/*
			create and invoke a function literal. Never have the different
			go routines working on the same variable as this will lead to
			strange results in the execution and will probably have some
			compilation warnings too.
		*/
		go func(link string) {
			time.Sleep(3 * time.Second)
			checkLink(link, c)
		}(l)
		// note we're invoking the function literal with the value from
		// the outer loop.
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Printf("❌ : %s \n", link)
		c <- link
		return
	}

	fmt.Printf("✅ : %s \n", link)
	c <- link
}
