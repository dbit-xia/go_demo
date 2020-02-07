package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/akrylysov/pogreb"
)

func main() {
	db, err := pogreb.Open("/tmp/pogreb.test", &pogreb.Options{BackgroundSyncInterval: 10 * time.Second})
	if err != nil {
		log.Fatal(err)
		return
	}

	count := 50 * 10000

	// fmt.Println(time.Now(), "Put start")

	// for index := 0; index < count; index++ {
	// 	err := db.Put([]byte("testKey"+strconv.FormatInt(int64(index), 10)), []byte("testValue"))
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	if index%10000 == 0 {
	// 		fmt.Println(time.Now(), "Put", index, err)
	// 	}
	// }
	// fmt.Println(time.Now(), "Put end")

	fmt.Println(time.Now(), "ReadItems start")
	for index := 0; index < count; index++ {
		key := []byte("testKey" + strconv.FormatInt(int64(index), 10))
		var val []byte

		// val, err := db.Get(key)
		err := db.Delete(key)
		if err != nil {
			log.Fatal(err)
		}
		if index%10000 == 0 {
			log.Println(index, string(key), string(val))
		}
	}

	// it := db.Items()
	// readCount := 0
	// for {
	// 	key, val, err := it.Next()
	// 	if err != nil {
	// 		if err != pogreb.ErrIterationDone {
	// 			log.Fatal(err)
	// 		}
	// 		break
	// 	}
	// 	readCount++
	// 	db.Delete(key)
	// 	if readCount%10000 == 0 {
	// 		log.Println(readCount, string(key), string(val))
	// 	}
	// }
	// db.Sync()
	fmt.Println(time.Now(), "ReadItems end")

	defer db.Close()
	var input string
	fmt.Scan(&input)
}
