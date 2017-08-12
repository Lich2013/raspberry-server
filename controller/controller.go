package controller

import (
	"github.com/gin-gonic/gin"
	"errors"
	"strconv"
)

type Base struct{}

func (this *Base) GetStringParam(c *gin.Context, key string) (string, error) {
	s, b := c.GetQuery(key)
	if b {
		return s, nil
	}
	s, b = c.GetPostForm(key)
	if b {
		return s, nil
	}
	return "", errors.New("this key doesn't exist")
}

func (this *Base) GetIntParam(c *gin.Context, key string) (value int, err error) {
	s, b := c.GetQuery(key)
	if !b {
		s, b = c.GetPostForm(key)
	}
	if !b {
		return 0, errors.New("this key doesn't exist")
	}
	value, err = strconv.Atoi(s)
	if err != nil {
		return 0, errors.New("this param is not int type")
	}
	return
}
