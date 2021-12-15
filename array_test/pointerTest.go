package main

import "fmt"

func main() {
	var a = []int{1, 2, 3}
	var b *[]int
	b = &a
	(*b)[0] = 2
	fmt.Println(a)
}
