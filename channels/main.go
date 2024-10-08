package main

import (
	"fmt"
	"net/http"
	"time"
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

	for link := range c {
		go func(l string) {
			time.Sleep(time.Second * 5)
			checkLink(l, c)
		}(link)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link + " might be down!")
		c <- link
		return
	}

	fmt.Println(link + " appears to be up!")
	c <- link
}
