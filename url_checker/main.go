package main

import (
	"errors"
	"fmt"
	"net/http"
)

type result struct {
	url    string
	status string
}

var errRequestFailed = errors.New("request failed")

func main() {

	results := make(map[string]string) // 초기화 하지 않으면 nil 값이 result에 들어가는 문제 있음

	c := make(chan result)
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

	for i := 0; i < len(urls); i++ {
		res := <-c
		results[res.url] = res.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, c chan<- result) {
	// chan<- 를 지정하면서 send-only로 만들 수 있음
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "Failed"
	}
	c <- result{url: url, status: status}
}
