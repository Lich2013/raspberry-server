package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"raspberry-server/receiver"
	"raspberry-server/pull"
)

func RegisterRouters() {
	auth := router.Group("/", Auth())
	auth.GET("index", func(context *gin.Context) {
		data := map[string]interface{}{"test": 123}
		context.JSON(http.StatusOK, data)
	})

	router.POST("receive", receiver.Receiver{}.Receive)
	router.GET("tasklist", pull.Pull{}.TaskList)
	router.POST("confirm", receiver.Receiver{}.Confirm)
}
