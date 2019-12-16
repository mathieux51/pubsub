package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v7"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {

	rdb := redis.NewClient(&redis.Options{
		Addr: ":6379",
	})

	pong, err := rdb.Ping().Result()
	if err != nil {
		return err
	}

	log.Println("redis", pong)

	pubsub := rdb.Subscribe("mychannel1")

	// Wait for confirmation that subscription is created before publishing anything.
	_, err = pubsub.Receive()
	if err != nil {
		return err
	}

	// Go channel which receives messages.
	ch := pubsub.Channel()

	// Publish a message.
	err = rdb.Publish("mychannel1", "hello").Err()
	if err != nil {
		return err
	}

	time.AfterFunc(time.Second, func() {
		// When pubsub is closed channel is closed too.
		_ = pubsub.Close()
	})

	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
	}

	return nil
}
