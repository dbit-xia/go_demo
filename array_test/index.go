package main

import (
	"fmt"
)

func main() {

	//var total int
	//total = 1 * 1000 * 1000
	//var slice = make([]string, total)
	//
	//fmt.Println(time.Now())
	//
	//for index := 0; index < total; index++ {
	//	slice[index] = fmt.Sprint("elem:", index)
	//	if index == total-1 {
	//		fmt.Println("EnqueueString")
	//	}
	//}
	//fmt.Println(time.Now())
	//for index := 0; index < total; index++ {
	//	elem := slice[index]
	//	if index == total-1 {
	//		fmt.Println("DequeueString=", elem)
	//	}
	//}
	//// fmt.Println(slice)
	//
	//// time.Sleep(time.Minute * 5)

	var a []int
	a = make([]int, 1, 10)
	//a[0] = 0
	//a[1] = 1
	//a[2] = 2
	//a[3] = 3
	a = append(a, 0)
	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3)
	var key = "xxx"
	var s = map[string]interface{}{key: 1}
	for index, value := range s {
		fmt.Println(index, value)
	}
	fmt.Println(a, len(a), cap(a), a[1:3])
}
