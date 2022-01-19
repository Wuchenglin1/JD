package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const user = "root:Wcl2021214174..@tcp(110.42.165.192:3306)/RedRock_WinterWork"

var dB *sql.DB

func InitMySql() {
	db, err := sql.Open("mysql", user)
	if err != nil {
		panic(err)
	}
	dB = db
}
