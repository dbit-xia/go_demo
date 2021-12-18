package main

import (
	"fmt"
	"time"
)

func main() {

	var m = make(map[int]string)

	type goResult struct {
		index int
		err   error
	}

	okCount := int32(0)
	var result = make(chan *goResult, 0)
	for i := 0; i < 100; i++ {
		go func(index int, result chan *goResult) {
			//result.index = index
			time.Sleep(time.Second)
			fmt.Println("send ", index)
			result <- &goResult{index: index}
		}(i, result)
	}

	for i := 0; i < 100; i++ {
		result2 := <-result
		m[result2.index] = "123"
		fmt.Println("receive", result2.index)
	}

	//count:=0
	//
	//	counter.m.Range(func(a interface{}, b interface{}) bool {
	//		count++
	//		//fmt.Println(count,a)
	//		return true
	//	})

	time.Sleep(2 * time.Second)
	fmt.Println(okCount, "", len(m))
}
