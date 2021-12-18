package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var m sync.Map

	type goResult struct {
		index int
		err   error
	}

	okCount := int32(0)
	var result = make(chan *goResult, 0)
	for index := 1; index <= 1000; index++ {
		go func(index int, result chan *goResult) {
			//m[index - 1] = index
			m.Store(index, index)
		}(index, result)
	}

	count := 0

	time.Sleep(1 * time.Second)
	//for _, value := range m {
	//	if value > 0 {
	//		count++
	//	}
	//}

	m.Range(func(key, value interface{}) bool {
		count++
		return true
	})

	fmt.Println(okCount, "", count)
}
