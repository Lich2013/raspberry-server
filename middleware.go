package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, bool := c.GetPostForm("token")
		if !bool || token != "asdf" {
			c.JSON(http.StatusOK, map[string]interface{}{"status": 401, "msg": "auth failed"})
			return
		}
	}
}
