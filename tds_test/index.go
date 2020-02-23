package main

import (
	"database/sql"
	"log"
	"strings"
	"fmt"

	_ "github.com/thda/tds"
)

func main() {
	cnxStr := "tds://zgsybase:zgsybase@192.168.15.107:51070/rts_doson?charset=utf8&encryptPassword=no"
	db, err := sql.Open("tds", cnxStr)
	// log.Println(db, err)
	if err != nil {
		log.Fatal(err)
	}

	// id := 2
	rows, err := db.Query("select outdate,nos from u2sale")
	if err != nil {
		log.Fatal(err)
	}

	var columns []string
	columns, err = rows.Columns()
	log.Println(" ", strings.Join(columns, ","))

	var maxCount = 10000000
	// var results = make([][]string, maxCount)
	var readCount = 0
	for rows.Next() {
		var dbno string
		var names string
		if err := rows.Scan(&dbno, &names); err != nil {
			log.Fatal(err)
		}
		readCount++
		//results[readCount-1] = []string{dbno, names}

		if readCount % 10000 == 0{
			log.Println(readCount, strings.Join([]string{dbno, names}, ","))	
		}
		
		if readCount >= maxCount {
			// break //都可以中断读取
			rows.Close() //都可以读取
		}
	}

	// for i := 1; i <= readCount; i++ {
	// 	log.Println(i, strings.Join(results[i-1], ","))
	// }

	var temp string
	fmt.Scan(&temp)

	db.Close()
}
