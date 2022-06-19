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

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*400)
	defer func() { cancel() }()

	asyncAll := gaw.AsyncAll[*http.Response](ctx,
		[]gaw.Function[*http.Response]{f1, f2, f3}...)

	// do other work while waiting async to finish
	fmt.Println("do other work")

	// call Await
	asyncAll.Await()

	fmt.Println("work done")

	for _, r := range asyncAll {

		if r.IsErr() {
			fmt.Println("err: ", r.Err())
			fmt.Println("----")
		} else {
			fmt.Println(r.Get())
			fmt.Println("----")
		}
	}
}
