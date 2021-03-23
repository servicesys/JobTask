package core

import (
	"fmt"
	"github.com/bamzi/jobrunner"
)

type TaskService struct {

	taskStorange TaskStorage
	taskTypes   []TaskType
}

func (taskService *TaskService) RegisterTaskJob(job TaskJob) {

	runner := Runner{
		taskJob:     job,
		taskStorage: taskService.taskStorange,
	}
	errorSchedule := jobrunner.Schedule("@every 5s",  runner)
	fmt.Println(errorSchedule)

}


func (taskService *TaskService) Start() {

	jobrunner.Start()
	//Obter todas as task type e registrar

}


