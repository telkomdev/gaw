## Gaw (Go Async Await)

The absurd implementation of `Node's Async Await` with `Golang`

### Requirements
- Go 1.18+ (because `gaw` uses generic features)

### Install

```shell
$ go get github.com/telkomdev/gaw
```

### Basic Usage

```go
package main

import (
	"context"
	"fmt"
	"github.com/telkomdev/gaw"
	"time"
)

type Event struct {
	ID   string
	Name string
}

func (e *Event) String() string {
	return fmt.Sprintf("event id: %s\nevent name: %s\n", e.ID, e.Name)
}

func main() {
	async := gaw.Async[*Event](context.Background(), func() (*Event, error) {
		// simulate heavy work that takes 3 seconds to finish
		time.Sleep(time.Second * 3)

		return &Event{ID: "111", Name: "Order Request"}, nil
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
        // no error
        // get the value
		fmt.Println(async.Get())
	}
}

```

### AsyncAll

Call multiple function in `async` mode, and get the multiple results

```go
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
```