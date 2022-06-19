package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/telkomdev/gaw"
)

func httpGet(url string) (*http.Response, error) {
	transport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
		IdleConnTimeout:     10 * time.Second,
	}

	httpClient := &http.Client{
		//Timeout:   time.Second * 10,
		Transport: transport,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	response, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func main() {
	url := "https://www.w3schools.com/js/default.asp"

	async := gaw.Async[*http.Response](context.Background(), func() (*http.Response, error) {

		resp, err := httpGet(url)

		return resp, err
	})

	// do other work while waiting async to finish
	fmt.Println("do other work")

	// call Await
	async.Await()

	fmt.Println("work done")

	// check if its error
	if async.IsErr() {
		fmt.Printf("async error: %v\n", async.Err())
	} else {
		resp := async.Get()
		fmt.Println(resp)
	}
}
