package main

import (
	"JD/api"
	"JD/dao"
)

func main() {
	dao.InitMySql()
	api.InitRouter()
}
