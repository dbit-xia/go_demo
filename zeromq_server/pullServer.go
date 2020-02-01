//
//  Hello World server.
//  Binds REP socket to tcp://*:5555
//  Expects "Hello" from client, replies with "World"
//

package main

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
)

func main() {
	//  Socket to talk to clients
	socket, _ := zmq.NewSocket(zmq.PULL)
	// defer responder.Close()
	// socket.SetConflate(true)
	socket.SetRcvbuf(100000)
	socket.SetRcvhwm(100000)
	// socket.SetConflate(false)
	socket.Bind("tcp://127.0.0.1:5555")
	

	var i int64
	for {
		//  Wait for next request from client
		msg, err := socket.Recv(0)
		// if err == nil {
		i++
		fmt.Println(i, "Received ", msg, err)
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
