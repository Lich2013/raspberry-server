package receiver

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"raspberry-server/controller"
	"raspberry-server/task"
	"crypto/sha256"
	"time"
	"encoding/hex"
	"strconv"
)

type Receiver struct {
	*controller.Base
}

func (this Receiver) Receive(context *gin.Context) {
	taskName, err := this.GetStringParam(context, "name")
	taskType, err := this.GetStringParam(context, "type")
	TaskDetail, err := this.GetStringParam(context, "detail")
	if err != nil {
		data := map[string]interface{}{"status": 403, "msg": err.Error()}
		context.JSON(http.StatusOK, data)
		return
	}
	h := sha256.New()
	str := taskName + strconv.FormatInt(time.Now().UnixNano(), 10)
	h.Write([]byte(str))
	taskId := hex.EncodeToString(h.Sum(nil))
	t := task.Task{
		TaskName:   taskName,
		TaskType:   taskType,
		TaskDetail: TaskDetail,
		TaskId:     taskId,
	}
	err = task.TaskAdd(&t)
	if err != nil {
		context.JSON(http.StatusOK, map[string]interface{}{"status": 500, "msg": err.Error()})
		return
	}
	context.JSON(http.StatusOK, map[string]interface{}{"status": 200, "msg": "ok"})
}

//todo dependent disk data
func (this Receiver) Confirm(context *gin.Context) {
	tasklist, err := this.GetStringParams(context, "tasklist") //todo 增加回调校验的参数
	if err != nil {
		data := map[string]interface{}{"status": 403, "msg": err.Error()}
		context.JSON(http.StatusOK, data)
		return
	}
	if len(*tasklist) == 0 {
		context.JSON(http.StatusOK, map[string]interface{}{"status": 404, "msg": "none"})
		return
	}
	if !task.TaskDel(tasklist) {
		context.JSON(http.StatusOK, map[string]interface{}{"status": 500, "msg": "length is incorrect"})
		return
	}
	context.JSON(http.StatusOK, map[string]interface{}{"status": 200, "msg": "ok"})
	return
}
