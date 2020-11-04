package utils

import (
	"database/sql"
	//数据库驱动,不使用也必须导入
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/go_stu")
	if err != nil {
		panic(err.Error())
	}
}
