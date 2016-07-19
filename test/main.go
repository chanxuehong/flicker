package main

import (
	"database/sql"
	"fmt"

	"github.com/chanxuehong/flicker"
	_ "github.com/go-sql-driver/mysql"
)

const DSN = "root:password@tcp(127.0.0.1:3306)/dbtest?clientFoundRows=false&parseTime=true&loc=Asia%2FShanghai&timeout=5s&collation=utf8mb4_general_ci"

var DB *sql.DB                   // 一般是全局变量, 不需要 DB.Close()
var Generator *flicker.Generator // 一般是全局变量, 不需要 Generator.Close()

func init() {
	var err error
	DB, err = sql.Open("mysql", DSN)
	if err != nil {
		panic(err.Error())
	}
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(20)

	Generator, err = flicker.NewGenerator(DB)
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(Generator.NextID())
	}
}
