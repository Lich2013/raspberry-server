package main

import (
	"github.com/gin-gonic/gin"
	"raspberry-server/task"
	"raspberry-server/conf"
)

var (
	router *gin.Engine
)

func main() {
	conf.LoadConfig()
	go task.TaskListen()
	router = gin.Default()
	router.Use()
	RegisterRouters()
	router.Run(conf.Conf.Host+":"+conf.Conf.Port)
}
