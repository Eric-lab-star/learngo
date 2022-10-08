package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestError = errors.New("request Failed")

func main() {
	var results = make(map[string]string)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}
	for _, url := range urls {
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "Failed"
		}
		results[url] = result
	}
	for key, value := range results {

		fmt.Println(key, value)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking url:", url)
	resp, err := http.Get(url)

	if err != nil || resp.StatusCode >= 404 {
		fmt.Println(resp.StatusCode)
		return errRequestError
	}

	return nil
}
