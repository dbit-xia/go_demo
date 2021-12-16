package main

import (
	"app/utils"
	"fmt"
)

//var shareNum uint32 = 0

func sum(no int, s []int) int {
	//fmt.Println(no)
	sum := 0
	for _, v := range s {
		sum += v
	}
	//if no >= 3 {
	//	var a = 0
	//	fmt.Println(1 / a)
	//}
	//var newNum = atomic.AddUint32(&shareNum, 1)
	//fmt.Println(newNum)
	//time.Sleep(time.Second)

	return no
}

func main() {

	var fns = make([]func() interface{}, 20)
	for i := 0; i < 20; i++ {
		i := i
		fns[i] = func() interface{} {
			return sum(i, []int{2, 3, 4})
		}
	}

	var results, errors = utils.ParallelLimit(&fns, 3)
	//time.Sleep(10 * time.Second)
	fmt.Println("OK", *results, errors)
}
