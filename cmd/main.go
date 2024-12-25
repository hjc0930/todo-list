package main

import (
	"todo-list/config"
	"todo-list/pkg/utils"
	"todo-list/repository/db/dao"
	"todo-list/router"
)

func main() {
	config.InitConfig()
	dao.MysqlInit()
	utils.InitLog()

	r := router.NewRouter()
	r.Run(config.Config.System.HttpPort) // listen and serve on 0.0.0.0:8080
}
