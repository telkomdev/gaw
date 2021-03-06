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
		fmt.Println(async.Get())
	}
}
