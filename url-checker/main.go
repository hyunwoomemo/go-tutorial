package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

var errReqeustFailed = errors.New("Request failed")

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

		statusCode, err := hitURL(url)
		if err != nil {
			result = strconv.Itoa(statusCode)
		}
		results[url] = result
	}

	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string) (statusCode int, error error) {

	fmt.Println("Checking", url)

	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		return resp.StatusCode, errReqeustFailed
	}

	return resp.StatusCode, nil
}

