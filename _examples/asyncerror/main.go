package main

import (
	"context"
	"errors"
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

		return nil, errors.New("request error")
	})

	fmt.Println("do other work")

	async.Await()

	fmt.Println("work done")

	if async.IsErr() {
		fmt.Printf("asyng error: %v\n", async.Err())
	} else {
		fmt.Println(async.Get())
	}
}
