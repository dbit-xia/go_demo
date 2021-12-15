package main

import (
	"fmt"
	"time"
)

func sum(no int, s []int) interface{} {
	//fmt.Println(no)
	sum := 0
	for _, v := range s {
		sum += v
	}
	time.Sleep(time.Second)
	return no
}

func parallelLimit(fns *[]func() interface{}, limit int) *[]interface{} {
	c := make(chan int)
	var runningCount = 0
	var total = len(*fns)
	var results = make([]interface{}, len(*fns))
	for i := 0; i < total; i++ {
		//fmt.Println("i", i)
		if runningCount < limit {
			runningCount++
		} else {
			var value = <-c
			fmt.Println(time.Now(), value)
		}

		i := i
		go func(c chan int) {
			results[i] = (*fns)[i]()
			c <- i
		}(c)
	}

	for i := 0; i < limit; i++ {
		var value = <-c
		fmt.Println(time.Now(), value)
	}
	close(c)

	return &results
}

func main() {

	var fns = make([]func() interface{}, 20)
	for i := 0; i < 20; i++ {
		i := i
		fns[i] = func() interface{} {
			return sum(i, []int{2, 3, 4})
		}
	}

	var results = parallelLimit(&fns, 10)

	fmt.Println("OK", *results)
}
