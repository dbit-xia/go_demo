package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

type any interface{}

//var shareNum uint32 = 0

func sum(no int, s []int) int {
	//fmt.Println(no)
	sum := 0
	for _, v := range s {
		sum += v
	}
	//if no >= 3 {
	//	var a = 0;
	//	fmt.Println(1 / a)
	//}
	//var newNum = atomic.AddUint32(&shareNum, 1)
	//fmt.Println(newNum)
	//time.Sleep(time.Second)

	return no
}

func parallelLimit(fns *[]func() any, limit int) (*[]any, *[]any) {
	c := make(chan int)
	var runningCount = 0
	var total = len(*fns)
	var resolves = make([]any, total)

	var errorMap = make(map[int]any, total)
	var lastErrorIndex uint32 = 0
	var hasError = false
	for index := 0; index < total; index++ {

		if runningCount < limit {
			runningCount++
		} else {
			var value = <-c
			err, _ := errorMap[value]
			fmt.Println(time.Now(), value, err)
			if hasError == true {
				runningCount-- //正在执行减1
				break
			}
		}

		index2 := index
		go func(c chan int) {
			defer func() {
				// 发生宕机时，获取panic传递的上下文并打印
				err := recover()
				switch err.(type) {
				case runtime.Error: // 运行时错误
					fmt.Println("runtime error:", err, index2)
					atomic.AddUint32(&lastErrorIndex, 1)
					errorMap[index2] = err
					hasError = true
				default: // 非运行时错误
					//fmt.Println("error:", err)
				}
				c <- index2
			}()

			resolves[index2] = (*fns)[index2]()

		}(c)
	}

	for i := 0; i < runningCount; i++ {
		var value = <-c
		fmt.Println(time.Now(), value, errorMap[value])
	}
	close(c)

	var errorArray []any
	if hasError {
		errorArray = make([]any, 0)

		for _, value := range errorMap {
			errorArray = append(errorArray, value)
		}
	}

	return &resolves, &errorArray
}

func main() {

	var fns = make([]func() any, 20)
	for i := 0; i < 20; i++ {
		i := i
		fns[i] = func() any {
			return sum(i, []int{2, 3, 4})
		}
	}

	var results, errors = parallelLimit(&fns, 3)
	//time.Sleep(10 * time.Second)
	fmt.Println("OK", *results, *errors)
}
