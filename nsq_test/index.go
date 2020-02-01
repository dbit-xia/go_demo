package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/nsqio/go-nsq"
)

func main() {
	var input string
	fmt.Println("生产消息(P)还是消费消息(C)?")
	fmt.Scan(&input)
	switch input {
	case "P":
		publish()
	case "C":
		subcribe()
	}
}

var received int

type myMessageHandler struct{}

// HandleMessage implements the Handler interface.
func (h *myMessageHandler) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		// Returning nil will automatically send a FIN command to NSQ to mark the message as processed.
		return nil
	}
	received++
	if received%10000 == 0 {
		fmt.Println(time.Now(), received, len(m.Body))
	}
	// err := processMessage(m.Body)

	// Returning a non-nil error will automatically send a REQ command to NSQ to re-queue the message.
	return nil
}

func subcribe() {
	const topicName = "topic"
	// Instantiate a consumer that will subscribe to the provided channel.
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(topicName, "channel", config)
	if err != nil {
		log.Fatal(err)
	}

	var input string
	fmt.Print("接收消息?(Y)")
	fmt.Scan(&input)
	if input == "Y" {
		// Set the Handler for messages received by this Consumer. Can be called multiple times.
		// See also AddConcurrentHandlers.
		consumer.AddHandler(&myMessageHandler{})
		// Use nsqlookupd to discover nsqd instances.
		// See also ConnectToNSQD, ConnectToNSQDs, ConnectToNSQLookupds.
		err = consumer.ConnectToNSQLookupds([]string{"127.0.0.1:4163", "127.0.0.1:4161"})
		if err != nil {
			log.Fatal("ConnectToNSQD", err)
		}
	}

	fmt.Print("退出?(Y)")
	fmt.Scan(&input)
	// Gracefully stop the consumer.
	if input == "Y" {
		consumer.Stop()
	}
}

func publish() {
	const topicName = "topic"

	// Instantiate a producer.
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer("127.0.0.1:4152", config)
	if err != nil {
		log.Fatal(err)
	}

	maxBatchCount := 10000
	totalSize := 0
	eleCount := 0
	text := []byte(strings.Repeat("hello", 200))
	messageBodys := make([][]byte, maxBatchCount)

push:
	for index := 0; index < 10000000; index++ {
		totalSize += len(text)
		eleCount++
		messageBodys[index] = text
		if totalSize >= 5*1000*1000 || eleCount >= maxBatchCount {
			break push
		}
	}
	messageBodys2 := messageBodys[0:eleCount]

	sendCount := 50 * 10000 / eleCount
	fmt.Println(eleCount, totalSize, len(messageBodys2), sendCount)
	var input string
	fmt.Println("批量发送?(Y)")
	fmt.Scan(&input)

	if input == "Y" {
		for index := 0; index < sendCount; index++ {

			// Synchronously publish a single message to the specified topic.
			// Messages can also be sent asynchronously and/or in batches.
			err = producer.MultiPublish(topicName, messageBodys2)
			if err != nil {
				log.Fatal(err)
			}
			if index%1 == 0 {
				fmt.Println(time.Now(), index)
			}
		}
		fmt.Println(time.Now(), "send OK")
	}

	fmt.Println(time.Now(), "Gracefully stop the producer?(Y)")
	fmt.Scan(&input)
	if input == "Y" {
		producer.Stop()
	}

	fmt.Println(time.Now(), "退出")
	fmt.Scan(&input)
}
