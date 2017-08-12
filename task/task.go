package task

import (
	"errors"
	"fmt"
)

var (
	TaskChan chan Task
)
const Tasklength = 1000

type Task struct {
	TaskName string `json:"task_name"`
	TaskType string `json:"task_type"`
	TaskDetail string `json:"task_detail"`
}

func TaskListen()  {
	TaskChan = make(chan Task, Tasklength)
}

func TaskAdd(task *Task) (err error) {
	if len(TaskChan) > 0 {
		tmpChan := TaskChan
		TaskChan = make(chan Task, len(TaskChan) + Tasklength) //fixme: 太蠢了
		for x := range tmpChan {
			TaskChan <- x
		}
	}
	if len(TaskChan) > Tasklength {
		return errors.New("this queue has full")
	}
	fmt.Println(task)
	TaskChan <- *task
	close(TaskChan)
	return nil
}

func TaskList() *[]Task {
	data := new([]Task)
	for x := range TaskChan {
		*data = append(*data, x)
	}
	fmt.Println(data)
	return data
}

//fixme: 删除任务时, 如果新增任务?
func TaskDel(length int) bool {
	if len(TaskChan) < length {
		return false
	}
	for i := 0; i < length; i++ {
		<-TaskChan
	}
	return true
}