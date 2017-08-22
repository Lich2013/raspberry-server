package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"raspberry-server/receiver"
	"raspberry-server/pull"
)

func RegisterRouters() {
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{})
	})
	auth := router.Group("/", Auth())
	auth.POST("receive", receiver.Receiver{}.Receive)
	auth.GET("tasklist", pull.Pull{}.TaskList)
	auth.POST("confirm", receiver.Receiver{}.Confirm)
}
