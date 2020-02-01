package main

import (
	"log"

	"github.com/zeromq/goczmq"
)

func main() {
	// Create a router socket and bind it to port 5555.
	router, err := goczmq.NewPull("tcp://*:5555")
	if err != nil {
		log.Fatal(err)
	}
	// defer router.Destroy()
	log.Println("router created and bound")

	var i int64
	for {
		buf := make([]byte, 16386)
		n, err := router.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		i++
		log.Println(i, "router receive", len(buf), n, err)

	}
}
