package main

import (
	"database/sql"
	"fmt"

	"github.com/chanxuehong/flicker"
	_ "github.com/go-sql-driver/mysql"
)

const DSN = "root:password@tcp(127.0.0.1:3306)/swift_nuochou_com?clientFoundRows=false&parseTime=true&loc=UTC&timeout=5s&collation=utf8mb4_general_ci"

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("mysql", DSN)
	if err != nil {
		panic(err.Error())
	}
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(20)

	flicker.Init(DB) // 初始化 flicker
}

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(flicker.NextID())
	}
}
