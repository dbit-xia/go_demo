package main

import "fmt"

func main() {
	var (
		name    string
		age     int
		married bool
	)
	fmt.Println("请输入姓名:")
	fmt.Scan(&name)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}
