package main

import (
	"github.com/gin-gonic/gin"
	"raspberry-server/task"
)

var (
	router *gin.Engine
)

func main() {
	go task.TaskListen()
	router = gin.Default()
	router.Use()
	RegisterRouters()
	router.Run("127.0.0.1:3000")
}
