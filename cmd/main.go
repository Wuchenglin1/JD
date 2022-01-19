package main

import (
	"JD/api"
	"JD/dao"
)

func main() {
	api.InitRouter()
	dao.InitMySql()
}
