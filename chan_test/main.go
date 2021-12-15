package main

import (
	"fmt"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c

}

func sum2(s []int, c *int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	*c = sum // 把 sum 发送到通道 c
}

func main() {
	//c := make(chan int)
	//go sum([]int{1, 2, 3}, c)
	//go sum([]int{2, 3}, c)
	//
	//x, y := <-c, <-c //接收最后两次
	//fmt.Println(x, y)

	var n int
	go sum2([]int{1, 2, 3}, &n)

	c := make(chan int)
	go sum([]int{1, 2, 3}, c)

	fmt.Println(n, <-c)
}
