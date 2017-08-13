package task

import (
	"errors"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

var (
	TaskChan chan Task
	flag     chan int
)

const Tasklength = 1000

type Task struct {
	TaskName   string `json:"task_name"`
	TaskType   string `json:"task_type"`
	TaskDetail string `json:"task_detail"`
	TaskId     string `json:"task_id"`
}

func TaskListen() {
	TaskChan = make(chan Task, Tasklength)
	flag = make(chan int)
	flag <- 1
}

func TaskAdd(task *Task) (err error) {
	if len(TaskChan) > Tasklength {
		return errors.New("this queue has full")
	}
	go func() {
		<-flag
		read, err := ioutil.ReadFile("tasklist")
		if err != nil {
			ioutil.WriteFile("tasklist", nil, 0644)
		}
		taskList := new([]Task)
		json.Unmarshal(read, taskList)
		*taskList = append(*taskList, *task)
		data, err := json.Marshal(taskList)
		ioutil.WriteFile("tasklist", data, 0644)
		TaskChan <- *task
		flag <- 1
	}()
	return nil
}

func TaskList() *[]Task {
	data := []Task{}
	length := len(TaskChan)
	for i := 0; i < length; i++ {
		data = append(data, <-TaskChan)
	}
	return &data
}

//fixme: 删除任务时, 如果新增任务? 一致性如何保证?
//done
func TaskDel(idList *[]string) bool {
	go func() {
		<-flag
		read, err := ioutil.ReadFile("tasklist")
		if err != nil {
			ioutil.WriteFile("tasklist", nil, 0644)
		}
		taskList := []Task{}
		json.Unmarshal(read, &taskList)
		for _, x := range *idList {
			for index, y := range taskList {
				if x == y.TaskId {
					taskList = append(taskList[:index], taskList[index+1:]...)
				}
			}
		}
		data, err := json.Marshal(taskList)
		ioutil.WriteFile("tasklist", data, 0644)
		flag <- 1
	}()
	return true
}

//func TaskWrite(content []byte) {
//	<-flag
//	taskList := new([]Task)
//
//	read, err := ioutil.ReadFile("tasklist")
//	if err != nil {
//		ioutil.WriteFile("tasklist", nil, 0644)
//	}
//	json.Unmarshal(read, taskList)
//
//	<-flag
//}
//
//func TaskRead() *[]Task {
//	<-flag
//	taskList := new([]Task)
//
//	read, err := ioutil.ReadFile("tasklist")
//	if err != nil {
//		ioutil.WriteFile("tasklist", nil, 0644)
//	}
//	json.Unmarshal(read, taskList)
//	flag <- 1
//	return taskList
//}
