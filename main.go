package main

import (
	"fmt"
	"net/http"
)

type urlStruct struct {
	url    string
	status int
}

func main() {
	results := make(map[string]int)
	channel := make(chan urlStruct)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	for _, url := range urls {
		go hitUrl(url, channel)

	}

	for i := 0; i < len(urls); i++ {
		result := <-channel
		results[result.url] = result.status
	}
	for url, staus := range results {
		fmt.Println(url, staus)
	}

}

func hitUrl(url string, c chan<- urlStruct) {
	resp, _ := http.Get(url)
	c <- urlStruct{url: url, status: resp.StatusCode}

}
