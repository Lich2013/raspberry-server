package pull

import (
	"raspberry-server/controller"
	"raspberry-server/task"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Pull struct {
	*controller.Base
}

func (this Pull) TaskList(context *gin.Context) {
	data := task.TaskList()
	ret := map[string]interface{}{"status": 200, "msg": "ok", "data": data}
	context.JSON(http.StatusOK, ret)
}
