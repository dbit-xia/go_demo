package main

import (
	"fmt"
	"log"
	"strconv"

	badger "github.com/dgraph-io/badger"
)

func main() {
	// Open the Badger database located in the /tmp/badger directory.
	// It will be created if it doesn't exist. //.WithInMemory(false)//.WithNumMemtables(2)
	db, err := badger.Open(badger.DefaultOptions("/tmp/badger").WithMaxTableSize(10*1024))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Opened")
	// db.DropAll()
	for tindex := 0; tindex < 150; tindex++ {
		txn := db.NewTransaction(true)

		count := 10000

		for index := 0; index < count; index++ {
			// Your code here…
			key := []byte("answer" + strconv.FormatInt(int64(index), 10))
			value := []byte("1234567890abcdefghijklmnopqrstuvwxyz")
			err := txn.Set(key, value)
			if index%10000 == 0 {
				log.Println(tindex, index, err)
			}
		}

		key := "answer1"
		item, err := txn.Get([]byte(key))
		// value := item.String()
		// log.Println(value)
		var valNot, valCopy []byte
		err = item.Value(func(val []byte) error {
			// This func with val would only be called if item.Value encounters no error.

			// Accessing val here is valid.
			fmt.Printf("The answer is: %s\n", val)

			// Copying or parsing val is valid.
			valCopy = append([]byte{}, val...)

			// Assigning val slice to another variable is NOT OK.
			valNot = val // Do not do this.
			return nil
		})

		log.Println("Get "+key+"=", "valCopy=", string(valCopy), "valNot=", string(valNot))

		// Commit the transaction and check for error.
		if err = txn.Commit(); err != nil {
			log.Fatal(err)
		}

		txn.Discard()
	}

	var input string

	log.Println("UpdateEnd")

	log.Println("按任意键删除部分数据")
	fmt.Scan(&input)
	db.DropPrefix([]byte("answer"))
	log.Println("DropPrefixEnd")

	fmt.Println("按任意键关闭db:")
	fmt.Scan(&input)
	db.Close()
	log.Println("end")

	fmt.Println("按任意键退出程序")
	fmt.Scan(&input)
}
