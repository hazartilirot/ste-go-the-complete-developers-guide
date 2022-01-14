package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	c := make(chan string) // creating a channel

	for _, link := range links {
		go checkLink(link, c)
	}

	/*
		for {
			go checkLink(<-c, c)
		}

		an alternative syntax: range c means that for must wait for a channel to provide another incoming link */

	for l := range c {
		func(link string) {
			time.Sleep(time.Second)
			go checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}
	fmt.Println(link, "is up")
	c <- link
}
