package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://google.com",
		"https://youtube.com",
		"https://stackoverflow.com",
		"https://linkedin.com",
		"https://golang.org",
	}

	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(l string, c chan string) {
	_, err := http.Get(l)
	if err != nil {
		//fmt.Println(l + " is down!")
		c <- l + " is down!"
		return
	}

	//fmt.Println(l + " is up!")
	c <- l + " is up!"
}
