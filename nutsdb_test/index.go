package main

import (
	"log"
	"fmt"

	"github.com/xujiajun/nutsdb"
)

func main() {
	opt := nutsdb.DefaultOptions
	opt.Dir = "/tmp/nutsdb" //这边数据库会自动创建这个目录文件
	db, err := nutsdb.Open(opt)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.View(
		func(tx *nutsdb.Tx) error {
			key := []byte("name1")
			bucket := "bucket1"
			if e, err := tx.Get(bucket, key); err != nil {
				return err
			} else {
				fmt.Println(string(e.Value)) // "val1-modify"
			}
			return nil
		}); err != nil {
			log.Println(err)
	}

	if err := db.Update(
		func(tx *nutsdb.Tx) error {
		key := []byte("name1")
		val := []byte("val1")
		bucket := "bucket1"
		if err := tx.Put(bucket, key, val, 0); err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	var input string

	fmt.Scan(&input)

	//...
}