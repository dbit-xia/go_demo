package main

import (
	"fmt"
	"time"
)

func main() {
	//var a = &(struct{ A string }{A: "1"})
	//var b = a
	//b.A = "2"
	//fmt.Println(a, b)

	var a = struct{ A string }{A: "1"}
	var b = &a
	b.A = "b"
	a = struct{ A string }{A: "a2"}
	fmt.Println(a, b)

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(1 * time.Second)
}
