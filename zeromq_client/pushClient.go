//
//  Hello World server.
//  Binds REP socket to tcp://*:5555
//  Expects "Hello" from client, replies with "World"
//

package main

import (
	"fmt"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	//  Socket to talk to clients
	responder, _ := zmq.NewSocket(zmq.PUSH)
	// responder.SetSndbuf(128 * 1024 * 1024)
	// responder.SetSndhwm(1000000000)
	// responder.SetConflate(false)
	// responder.SetImmediate(true)
	// responder.SetReqRelaxed(100000)
	// defer responder.Close()
	responder.Connect("tcp://127.0.0.1:5555")

	// fmt.Scan(1)
	// responder.SendMessageDontwait()
	var i int64
	for {
		//  Wait for next request from client

		result, err := responder.SendMessage("123")
		i++
		fmt.Println(i, result, err)

		if i >= 10000 {
			break
		}
		// time.Sleep(time.Millisecond * 1)
		// } else {
		// 	time.Sleep(time.Second)
		// }

		//  Do some 'work'
		// time.Sleep(time.Second)

		//  Send reply back to client
		// reply := "World"
		// responder.Send(reply, 0)
		// fmt.Println("Sent ", reply)
	}
}
