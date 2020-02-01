package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	stan "github.com/nats-io/stan.go"
)

func main() {
	clusterID := "test-cluster"
	clientID := "test"
	sc, err := stan.Connect(clusterID, clientID)
	fmt.Println("Connect", err)
	// Simple Synchronous Publisher

	var input string
	fmt.Println("生产(P)还是消费(C)?")
	fmt.Scan(&input)

	if strings.Contains(input, "C") {
		receiveCount := 0
		// Simple Async Subscriber
		sub, err := sc.Subscribe("foo", func(m *stan.Msg) {
			receiveCount++
			if receiveCount%10000 == 0 {
				fmt.Println(time.Now().Format(time.StampMilli), receiveCount, "Received a message: ", string(m.Data))
			}
		}, stan.DurableName("my-durable"))
		fmt.Println("Subscribe", sub == nil, err)
	}

	if strings.Contains(input, "P") {
		count := 100 * 10000
		publishedCount := 0

		fmt.Println(time.Now(), count)
		value := []byte(strings.Repeat("Hello", 200))
		for index := 0; index < count; index++ {
			sc.PublishAsync("foo", value, func(lguid string, err error) {
				publishedCount++
				// glock.Lock()
				if publishedCount%10000 == 0 {
					log.Println(time.Now().Format(time.StampMilli), publishedCount, "published ACK for guid:", lguid)
				}

			})
			// err = sc.Publish("foo", []byte("Hello World")) // does not return until an ack has been received from NATS Streaming
			// if index%10000 == 0 {
			// 	fmt.Println(time.Now().Format(time.StampMilli), index, "Publish", err)
			// }
		}
	}

	// Unsubscribe
	// sub.Unsubscribe()

	fmt.Println("即将退出")
	fmt.Scan(&input)
	// Close connection
	sc.Close()
}
