package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"raspberry-server/conf"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := []string{}
		if values, _ := c.Request.Header["Token"]; len(values) > 0 {
			token = values
		} else {
			c.JSON(http.StatusOK, map[string]interface{}{"status": 401, "msg": "auth failed"})
			c.Abort()
			return
		}
		if token[0] != conf.Conf.Token {
			c.JSON(http.StatusOK, map[string]interface{}{"status": 401, "msg": "auth failed"})
			c.Abort()
			return
		}
		c.Next()

	}
}
