package main

import (
	"app/utils"
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"
	"time"
)

var iCnt int = 0

func helloHandler(w http.ResponseWriter, r *http.Request) {
	iCnt++
	str := "Hello world ! friend(" + strconv.Itoa(iCnt) + ")"
	io.WriteString(w, str)
	fmt.Println(str)
}

func main() {
	fmt.Println("main")
	ht := http.HandlerFunc(helloHandler)
	if ht != nil {
		http.Handle("/hello", ht)
	}

	//var waitGroup sync.WaitGroup

	var srv http.Server
	srv = http.Server{
		Addr: ":8080",
		BaseContext: func(listener net.Listener) context.Context {
			fmt.Println("listen success", srv.Addr)
			return context.TODO()
		},
	}
	var isExit bool

	//waitGroup.Add(1)
	var fns = []func() (interface{}, error){
		func() (interface{}, error) {
			//defer waitGroup.Done()
			fmt.Println("ListenAndServe")
			err := srv.ListenAndServe()
			if err != nil {
				strExit := strconv.FormatBool(isExit)
				fmt.Printf("ListenAndServe:%s isExit: %s \n", err.Error(), strExit)
			}
			return nil, err
		},
		func() (interface{}, error) {
			time.Sleep(3 * time.Second)
			isExit = true
			err := srv.Shutdown(context.Background())
			if err != nil {
				return nil, err
			}
			return nil, nil
		},
	}

	var wait = make(chan int, 0)
	go func() {
		defer func() {
			wait <- 0
		}()

		fmt.Println("wait")
		_, err := utils.ParallelLimit(&fns, 2)
		if err != nil {
			fmt.Println("ParallelLimit", err)
		}

		fmt.Println("end")
	}()

	fmt.Println("exit", <-wait)
}
