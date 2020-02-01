package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/nsqio/go-nsq"
)

func main() {
	// Instantiate a producer.
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
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

	sendCount := 150 * 10000 / eleCount
	fmt.Println(eleCount, totalSize, len(messageBodys2), sendCount)
	var input string
	fmt.Println("批量发送?(Y)")
	fmt.Scan(&input)

	if input == "Y" {
		for index := 0; index < sendCount; index++ {
			topicName := "topic"

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
