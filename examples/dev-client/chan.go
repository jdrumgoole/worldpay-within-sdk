package main

import (
	"fmt"
	"net/http"
	"time"
)

var urls = []string{
	"http://pulsoconf.co/",
	"http://golang.org/",
	"http://matt.aimonetti.net/",
}

type HttpResponse struct {
	url      string
	response *http.Response
	err      error
}

func testChannel() {

	//	results := asyncHttpGets(urls)
	//	for _, result := range results {
	//		fmt.Printf("%s status: %s\n", result.url,
	//			result.response.Status)
	//	}

	test2()
}

func asyncHttpGets(urls []string) []*HttpResponse {
	ch := make(chan *HttpResponse, len(urls)) // buffered
	responses := []*HttpResponse{}
	for _, url := range urls {
		go func(url string) {
			fmt.Printf("Fetching %s \n", url)
			resp, err := http.Get(url)
			resp.Body.Close()
			ch <- &HttpResponse{url, resp, err}
		}(url)
	}

	for {
		select {
		case r := <-ch:
			fmt.Printf("%s was fetched\n", r.url)
			responses = append(responses, r)
			if len(responses) == len(urls) {
				return responses
			}
		case <-time.After(50 * time.Millisecond):
			fmt.Printf(".")
		}
	}

	return responses
}

func test2() {

	cMsg := make(chan string)

	go func() {
		time.Sleep(1000 * time.Millisecond)
		cMsg <- "ping."
	}()

	msg := <-cMsg
	fmt.Println(msg)
}