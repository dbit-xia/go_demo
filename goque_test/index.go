package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/beeker1121/goque"
)

func main() {
	testQueue()
}

func testQueue() {
	var input string
	s, err := goque.OpenQueue("/tmp/goque/queue")
	fmt.Println(s, err)
	count := 150 * 10000

enqueue:
	fmt.Println("EnqueueString?(Y)")
	fmt.Scan(&input)
	if input == "Y" {
		for index := 0; index < count; index++ {
			item, err := s.EnqueueString(strings.Repeat("item value", 100))
			if index%10000 == 0 {
				fmt.Println(time.Now(), index, (item == nil), err)
			}
		}
	}

	fmt.Println("Dequeue?(Y)")
	fmt.Scan(&input)
	if input == "Y" {
		for index := 0; index < count; index++ {
			item, err := s.Dequeue()
			if index%10000 == 0 {
				if err != nil {
					fmt.Println(time.Now(), index, err)
				} else {
					fmt.Println(time.Now(), index, (item.ToString()))
				}

			}
		}
	}

	fmt.Println("enqueue too?(Y)")
	fmt.Scan(&input)
	if input == "Y" {
		goto enqueue
	}

	fmt.Println("删除数据?(Y)")
	fmt.Scan(&input)
	if input == "Y" {
		s.Drop()
	}

	fmt.Println("关闭?(Y)")
	fmt.Scan(&input)
	if input == "Y" {
		s.Close()
	}
	// ...
	fmt.Println("退出")
	fmt.Scan(&input)
}

func testStack() {
	var input string
	s, err := goque.OpenStack("/tmp/goque/stack")
	fmt.Println(s, err)
	count := 150 * 10000
	for index := 0; index < count; index++ {
		item, err := s.PushString("item value")
		if index%10000 == 0 {
			fmt.Println(time.Now(), index, item, err)
		}
	}
	// ...
	s.Close()

	fmt.Println("结束")
	fmt.Scan(&input)
}
