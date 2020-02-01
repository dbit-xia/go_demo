package main

import (
	"fmt"
	"time"

	"github.com/grandecola/bigqueue"
)

func main() {
	fmt.Println([]byte("elem"))
	//,bigqueue.SetPeriodicFlushOps(100000)
	var bqDir string
	bqDir = "/tmp/bigqueue"

	//SetArenaSize 每个文件大小
	//SetMaxInMemArenas //内存=几个文件大小之和, Close时才会释放内存
	//SetPeriodicFlushOps //写入多少个元素时 就刷入磁盘,值越大执行越快

	bq, err := bigqueue.NewMmapQueue(bqDir, bigqueue.SetArenaSize(10*1024*1024)) //, bigqueue.SetMaxInMemArenas(3)) //,bigqueue.SetPeriodicFlushDuration(time.Second))
	consumer, err := bq.NewConsumer("test")
	fmt.Println(bq, consumer, err)

	// defer bq.Close()
	var total int
	total = 15000000 / 10
	fmt.Println(time.Now())
	var elem string
	fmt.Println(elem, total)

	for index := 0; index < total; index++ {
		err = bq.Enqueue([]byte("elem"))
		if index%100000 == 0 {
			fmt.Println(index, "EnqueueString=", elem, err)
		}
	}
	fmt.Println(time.Now())

	for index := 0; index < total; index++ {
		elem, err = bq.DequeueString()
		// elem, err = consumer.DequeueString()
		if index%100000 == 0 {
			fmt.Println(index, "DequeueString=", elem, err)
		}
	}
	fmt.Println(time.Now())

	// J:
	// for index := 0; index < total; index++ {
	// 	if bq.IsEmpty()==true {
	// 		break J
	// 	}
	// 	elem, err = bq.DequeueString()
	// 	if index%100000 == 0 {
	// 		fmt.Println(index,"DequeueString=", elem, err)
	// 	}
	// }
	fmt.Println("IsEmpty=", bq.IsEmpty())
	// fmt.Println(bq)
	// bq.Flush()
	defer bq.Close() //

	// const cFilePerm = 0744
	// fd, err := os.OpenFile(bqDir+"/arena_0.dat", os.O_RDWR, cFilePerm)
	// m, err := mmap.NewSharedFileMmap(fd, 0, 128*1024*1024, 3)
	// err = m.Unmap()

	// fmt.Println(time.Now(), "Unmap", bq.am.arenas[0].Unmap())

	// time.Sleep(time.Minute * 100)
}
