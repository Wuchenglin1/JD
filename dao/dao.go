package dao

import (
	"JD/tool"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var dB *sql.DB

func InitMySql() {
	config := tool.GetConfig()
	db, err := sql.Open("mysql", config.MySql.User)
	if err != nil {
		panic(err)
	}
	dB = db
}
