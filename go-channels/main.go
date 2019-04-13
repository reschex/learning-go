package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(s string, c chan string) {
	_, err := http.Get(s)
	if err != nil {
		fmt.Println(s, "is down")
		c <- "Might be down I think"
	}
	fmt.Println(s, "is up")
	c <- "Yep, its up"
}
