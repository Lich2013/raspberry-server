package receiver

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"raspberry-server/controller"
)

type Receiver struct {
	*controller.Base
}

func (this Receiver) Receive(context *gin.Context) {
	id, err := this.GetIntParam(context, "id")
	if err != nil {
		data := map[string]interface{}{"status": 403, "msg": err.Error()}
		context.JSON(http.StatusOK, data)
		return
	}
	data := map[string]interface{}{"value": id}
	context.JSON(http.StatusOK, data)
}
