package utils

import (
	"fmt"
)

//func WrapError(wrapMsg string, err error) error {
//	pc, file, line, ok := runtime.Caller(4) //向前取4级堆栈位置
//	f := runtime.FuncForPC(pc)
//	if !ok {
//		return errors.New("WrapError 方法获取堆栈失败")
//	}
//	if err == nil {
//		return nil
//	} else {
//		errMsg := fmt.Sprintf("%s \n\tat %s:%d (Method %s)\nCause by: %s\n", wrapMsg, file, line, f.Name(), err.Error())
//		return errors.New(errMsg)
//	}
//}

//type any interface{}

type ParallelError struct {
	ErrorIndexes []int
	ErrorMap     map[int]error
}

type GoResult struct {
	index int
	err   error
}

func (p *ParallelError) Error() string {
	var msg = ""
	for key, err := range p.ErrorMap {
		msg += fmt.Sprintf("%d: %+v \n", key, err)
	}
	return msg
}

func ParallelLimit(fns *[]func() (interface{}, error), limit int) (*[]interface{}, error) {

	var runningCount = 0
	var total = len(*fns)
	var resolves = make([]interface{}, total)

	var parallelErrors = &ParallelError{
		//ErrorIndexes: []int{},
		ErrorMap: make(map[int]error, total),
	}
	//var lastErrorIndex uint32 = 0
	var errCount = 0

	c := make(chan *GoResult)
	var errors4 = &Errors{Skip: 4}

	for index := 0; index < total; index++ {
		//time.Sleep(time.Millisecond)
		if runningCount < limit {
			runningCount++
		} else {
			goResult := <-c
			if goResult.err != nil {
				errCount++
				parallelErrors.ErrorMap[goResult.index] = goResult.err

				runningCount-- //正在执行减1
				break
			}
		}

		go func(index int) {
			var err interface{}
			defer func() {
				// 发生宕机时，获取panic传递的上下文并打印
				if err == nil {
					err = recover()
					if err != nil {
						c <- &GoResult{index: index, err: errors4.WithStack(err.(error))}
						return
					}
				} else {
					c <- &GoResult{index: index, err: err.(error)}
					return
				}
				c <- &GoResult{index: index}
				return
			}()

			resolves[index], err = (*fns)[index]()

		}(index)
	}

	for i := 0; i < runningCount; i++ {
		goResult := <-c
		if goResult.err != nil {
			errCount++
			parallelErrors.ErrorMap[goResult.index] = goResult.err
		}
	}
	close(c)

	if errCount > 0 {
		parallelErrors.ErrorIndexes = make([]int, errCount)
		var i int32 = 0
		for key, _ := range parallelErrors.ErrorMap {
			parallelErrors.ErrorIndexes[i] = key
			i++
		}

		return &resolves, parallelErrors
	}

	return &resolves, nil
}
