package main

import (
	"app/utils"
	"fmt"
)

//var shareNum uint32 = 0

type Aaa struct {
	Name string
}

func sum(no uint, s []uint) (uint, error) {
	//fmt.Println(no)
	var sum uint
	sum = 0
	for _, v := range s {
		sum += v
	}
	if no >= 3 {
		var a = 0
		fmt.Println(1 / a)
	}
	//var newNum = atomic.AddUint32(&shareNum, 1)
	//fmt.Println(newNum)
	//time.Sleep(time.Second)
	//var a = &Aaa{Name: "张三"}
	return no, nil
}

//func wrap(fn interface{}, args ...interface{}) func() (interface{}, error) {
//	return func() (interface{}, error) {
//		return fn.(func(args ...interface{}) (interface{}, error))(args...)
//	}
//}

//func getType[T any](a interface{}) T {
//	var b T
//	b = a.(T)
//	return b
//}

func main() {

	var fns = make([]func() (interface{}, error), 20)
	for i := 0; i < 20; i++ {
		i := i
		fns[i] = func() (interface{}, error) {
			return sum(uint(i), []uint{2, 3, 4})
		}
		//fns[i] = func() (interface{}, error) {
		//	return utils.Call(sum, uint(i), []uint{2, 3, 4})
		//}
		//fns[i] = utils.Wrap(sum, uint(i), []uint{2, 3, 4})
	}

	var results, errors = utils.ParallelLimit(&fns, 3)
	//var msg=fmt.Sprintf("%+v",errors.(*utils.ParallelError).ErrorMap[3])
	fmt.Println("OK", *results, errors)
	//os.Stat("123")
}
