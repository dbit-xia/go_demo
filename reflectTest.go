package main

import (
	"app/utils"
	"fmt"
)

// 普通函数
func add(a, b int) {
	fmt.Println(a, b)
}

func main() {
	// 反射调用函数
	resolve, err := utils.Call(add, 1, 3)
	// 获取第一个返回值, 取整数值
	fmt.Println(resolve, err)
}
