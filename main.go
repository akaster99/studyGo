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
	results := make(map[string]string)
	for _, url := range urls {
		error := hitURL(url)
		if error != nil {
			results[url] = "fail"
		} else {
			results[url] = "success"
		}
	}
	for url, result := range results {
		fmt.Println(url, ": ", result)
	}
}

func hitURL(url string) error {
	resp, error := http.Get(url)
	if error != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
