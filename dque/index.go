package main

import (
	"fmt"
	"log"
	"time"

	"github.com/joncrlsn/dque"
)

// Item is what we'll be storing in the queue.  It can be any struct
// as long as the fields you want stored are public.
type Item struct {
	Name string
	ID   int
}

// ItemBuilder creates a new item and returns a pointer to it.
// This is used when we load a segment of the queue from disk.
func ItemBuilder() interface{} {
	return &Item{}
}

func main() {
	exampleDQueMain()
}

func exampleDQueMain() {
	qName := "item-queue"
	qDir := "/tmp"
	segmentSize := 100 * 1024 * 1024

	// Create a new queue with segment size of 50
	// q, err := dque.New(qName, qDir, segmentSize, ItemBuilder)

	// You can reconsitute the queue from disk at any time
	// as long as you never use the old instance
	q, err := dque.Open(qName, qDir, segmentSize, ItemBuilder)
	q.TurboOn()
	var total int
	total = 1000000
	fmt.Println(time.Now())
	// for index := 0; index < total; index++ {
	// 	err = q.Enqueue(&Item{"Joe", index})
	// 	if err != nil {
	// 		fmt.Println(index, err)
	// 	}
	// }

	fmt.Println(time.Now(), "Enqueue")

	// Peek at the next item in the queue
	var iface interface{}

	for index := 0; index < total; index++ {
		// if iface, err = q.Peek(); err != nil {
		// 	if err != dque.ErrEmpty {
		// 		log.Fatal("Error peeking at item ", err)
		// 	}
		// }
		// Dequeue the next item in the queue
		iface, err = q.Dequeue()

		log.Println(index, "dequeuing item ", err)
	}

	log.Println(time.Now(), "Dequeue", iface)
	time.Sleep(time.Hour)

	// Assert type of the response to an Item pointer so we can work with it
	item, ok := iface.(*Item)
	if !ok {
		log.Fatal("Dequeued object is not an Item pointer")
	}

	doSomething(item)
	time.Sleep(time.Hour)
}

func doSomething(item *Item) {
	log.Println("Dequeued", item)
}
