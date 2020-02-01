package main

import (
	"fmt"
	"time"
)

func main() {

	var total int
	total = 10000000
	var queueArray = make([]string, total)

	fmt.Println(time.Now())

	for index := 0; index < total; index++ {
		queueArray[index] = "elem" + string(index)
		if index == total-1 {
			fmt.Println("EnqueueString")
		}
	}
	fmt.Println(time.Now())
	for index := 0; index < total; index++ {
		elem := queueArray[index]
		if index == total-1 {
			fmt.Println("DequeueString=", elem)
		}
	}
	// fmt.Println(queueArray)

	time.Sleep(time.Minute * 5)
}
