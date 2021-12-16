package utils

import (
	"runtime"
	"strconv"
)

//type any interface{}

type ParallelError struct {
	errorIndexes []int
	errorMap     map[int]error
}

func (p *ParallelError) Error() string {
	var msg = ""
	for key, err := range p.errorMap {
		msg += strconv.Itoa(key) + " " + err.Error() + "\n"
	}
	return msg
}

func ParallelLimit(fns *[]func() (interface{}, error), limit int) (*[]interface{}, error) {
	c := make(chan int)
	var runningCount = 0
	var total = len(*fns)
	var resolves = make([]interface{}, total)

	var errors = &ParallelError{
		//errorIndexes: []int{},
		errorMap: make(map[int]error, total),
	}
	//var lastErrorIndex uint32 = 0
	var hasError = false
	for index := 0; index < total; index++ {
		//time.Sleep(time.Millisecond)
		if runningCount < limit {
			runningCount++
		} else {
			<-c
			//err, _ := errors.errorMap[okIndex]
			//fmt.Println(time.Now(), okIndex, err)
			if hasError == true {
				runningCount-- //正在执行减1
				break
			}
		}

		currentIndex := index
		go func(c chan int) {
			var err interface{}
			defer func() {
				// 发生宕机时，获取panic传递的上下文并打印
				if err == nil {
					err = recover()
				}
				if err != nil {
					hasError = true

					switch err.(type) {
					case runtime.Error: // 运行时错误
						errors.errorMap[currentIndex] = err.(runtime.Error)
					default: // 非运行时错误
						errors.errorMap[currentIndex] = err.(error)
					}
				}

				c <- currentIndex
			}()

			resolves[currentIndex], err = (*fns)[currentIndex]()

		}(c)
	}

	for i := 0; i < runningCount; i++ {
		<-c
		//fmt.Println(time.Now(), okIndex, errors.errorMap[okIndex])
	}
	close(c)

	if hasError {
		errors.errorIndexes = make([]int, len(errors.errorMap))
		var i int32 = 0
		for key, _ := range errors.errorMap {
			errors.errorIndexes[i] = key
			i++
		}

		return &resolves, errors
	}

	return &resolves, nil
}
