package main

import (
	"fmt"
	"reflect"
)

// 普通函数
func add(a, b int) (int, int) {
	return a, b
}

func call(fn interface{}, args ...interface{}) []interface{} {
	// 将函数包装为反射值对象
	funcValue := reflect.ValueOf(fn)
	// 构造函数参数, 传入两个整型值
	var paramList = make([]reflect.Value, len(args))
	for index, value := range args {
		paramList[index] = reflect.ValueOf(value)
	}

	// 反射调用函数
	retList := funcValue.Call(paramList)

	var result = make([]interface{}, len(retList))
	for index, value := range retList {
		v := value.Interface()
		result[index] = v
	}
	return result

}
func main() {
	// 反射调用函数
	retList := call(add, 1, 3)
	// 获取第一个返回值, 取整数值
	fmt.Println(retList)
}
