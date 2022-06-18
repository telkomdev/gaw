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

func main() {
	async := gaw.Async[Event](context.Background(), func() (Event, error) {
		// simulate heavy work that takes 3 seconds to finish
		time.Sleep(time.Second * 3)

		return Event{ID: "111", Name: "Order Request"}, nil
	})

	fmt.Println("do other work")

	fmt.Println("work done, maybe...")

	fmt.Println("result is empty, because the Await is emited: ", async.Get())
}
