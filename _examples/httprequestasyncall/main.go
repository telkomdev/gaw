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

	f1 := func() (*http.Response, error) {
		resp, err := httpGet(url)

		return resp, err
	}

	f2 := func() (*http.Response, error) {
		resp, err := httpGet(url)

		return resp, err
	}

	f3 := func() (*http.Response, error) {
		resp, err := httpGet(url)

		return resp, err
	}

	asyncAll := gaw.AsyncAll[*http.Response](context.Background(),
		[]gaw.Function[*http.Response]{f1, f2, f3}...)

	// do other work while waiting async to finish
	fmt.Println("do other work")

	// call Await
	asyncAll.Await()

	fmt.Println("work done")

	for _, resp := range asyncAll.Get() {

		fmt.Println(resp)
		fmt.Println("----")
	}
}
