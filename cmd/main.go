package main

import (
	"todo-list/config"
	"todo-list/repository/db/dao"
	"todo-list/router"
)

func main() {
	config.InitConfig()
	dao.MysqlInit()
	r := router.NewRouter()
	_ = r.Run(config.Config.System.HttpPort)

}
