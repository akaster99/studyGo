package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request Failed")

func main() {
	urls := []string{
		"https://www.airbnb.com",
		"https://www.google.com",
		"https://www.amazon.com",
		"https://www.reddit.com",
		"https://www.google.com",
		"https://soundcloud.com",
		"https://www.facebook.com",
		"https://www.instargam.com",
		"https://academy.nomadcoder.co/",
	}

	for _, url := range urls {
		hitURL(url)
	}
}

func hitURL(url string) error {
	fmt.Println("checking", url)
	resp, error := http.Get(url)
	if error != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
