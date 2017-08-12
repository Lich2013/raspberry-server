package main

import "github.com/gin-gonic/gin"
var (
	router *gin.Engine
)


func main() {
	router = gin.Default()
	RegisterRouters()
	router.Run("127.0.0.1:3000")
}