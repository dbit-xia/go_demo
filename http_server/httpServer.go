package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
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

	var waitGroup sync.WaitGroup

	var srv http.Server
	srv = http.Server{Addr: ":8080"}
	var isExit bool

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		fmt.Println("ListenAndServe")
		err := srv.ListenAndServe()
		if err != nil {
			strExit := strconv.FormatBool(isExit)
			fmt.Printf("ListenAndServe:%s isExit: %s \n", err.Error(), strExit)
		}

	}()

	go func() {
		time.Sleep(3 * time.Second)
		isExit = true
		err := srv.Shutdown(context.Background())
		if err != nil {
			return
		}
	}()

	fmt.Println("wait")
	waitGroup.Done()
	waitGroup.Wait()
	fmt.Println("end")
}
