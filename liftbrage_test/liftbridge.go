package main

import (
	"fmt"
	"time"

	lift "github.com/liftbridge-io/go-liftbridge"
	"golang.org/x/net/context"
)

func main() {
	// Create Liftbridge client.
	addrs := []string{"localhost:9292", "localhost:9293", "localhost:9294"}
	client, err := lift.Connect(addrs)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// Create a stream attached to the NATS subject "foo".
	var (
		subject = "foo"
		name    = "foo-stream"
	)
	if err := client.CreateStream(context.Background(), subject, name); err != nil {
		if err != lift.ErrStreamExists {
			panic(err)
		}
	}

	// Subscribe to the stream starting from the beginning.

	fmt.Println("Subscribe")
	ctx := context.Background()
	if err := client.Subscribe(ctx, name, func(msg *lift.Message, err error) {
		if err != nil {
			panic(err)
		}
		index := msg.Offset()
		if index%1 == 0 {
			fmt.Println(index, string(msg.Value()))
		}

		time.Sleep(1 * time.Second)

	}, lift.StartAtEarliestReceived()); err != nil {
		panic(err)
	}

	go func() {
		for i := 0; i < 1000; i++ {
			time.Sleep(1 * time.Second)
			fmt.Println("go func", i)
		}
	}()

	fmt.Println("Publish")
	for i := 0; i < 20; i++ {
		// Publish a message to "foo".
		if _, err := client.Publish(context.Background(), name, []byte("hello")); err != nil {
			panic(err)
		}
		if i%1 == 0 {
			fmt.Println(i)
		}
	}

	<-ctx.Done()

}
