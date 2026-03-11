package main

import (
	"fmt"
	"net/http"
)

type requestResult struct {
	URL string
	Success bool
	StatusCode int
}

func main() {

	results := make(map[string]int)
	c := make(chan requestResult)


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

		go hitURL(url, c)
	}

	for i:=0;i<len(urls);i++ {
		result := <-c

		results[result.URL] = result.StatusCode
	}


	fmt.Println(results)
}


func hitURL(url string, c chan<- requestResult) {

	resp, err := http.Get(url)

	success := true

	if err != nil || resp.StatusCode >= 400 {
		success = false
	} 
	
	c <-requestResult{URL: url, StatusCode: resp.StatusCode, Success: success}
}