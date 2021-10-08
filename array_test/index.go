package main

import (
	"fmt"
	"time"
)

func main() {

	var total int
	total = 1 * 1000 * 1000
	var queueArray = make([]string, total)

	fmt.Println(time.Now())

	for index := 0; index < total; index++ {
		queueArray[index] = fmt.Sprint("elem:", index)
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

	// time.Sleep(time.Minute * 5)
}
