package utils

import (
	"reflect"
)

func Call(fn interface{}, args ...interface{}) (interface{}, error) {
	// 将函数包装为反射值对象
	funcValue := reflect.ValueOf(fn)
	// 构造函数参数, 传入两个整型值
	var paramList = make([]reflect.Value, len(args))
	for index, value := range args {
		paramList[index] = reflect.ValueOf(value)
	}

	// 反射调用函数
	retList := funcValue.Call(paramList)
	var retCount = len(retList)

	if retCount == 0 {
		return nil, nil
	}

	var result = make([]interface{}, retCount)
	for index, value := range retList {
		v := value.Interface()
		result[index] = v
	}
	if retCount == 1 {
		return nil, result[0].(error)
	}

	if result[1] != nil {
		return result[0], result[1].(error)
	}

	return result[0], nil
}
