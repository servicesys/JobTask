package core

import (
	"fmt"
	"github.com/bamzi/jobrunner"
)

type TaskService struct {
	TaskStorage TaskStorage
	taskTypes   map[string]*TaskType
}

func (taskService *TaskService) RegisterTaskJob(taskJob TaskJob) {

	keyTaskJob := taskJob.GetTaskTypeName()
	fmt.Println(keyTaskJob)
	if taskType, found := taskService.taskTypes[keyTaskJob]; found {
		taskType.TaskJobRef = taskJob
	} else {
		panic("NOT EXIST TASK TYPE WITH NAME: [" + keyTaskJob + "]")
	}
}

func (taskService *TaskService) Start() {

	jobrunner.Start()

	for _, taskType := range taskService.taskTypes {

		if taskType.TaskJobRef == nil {
			panic("TaskJob NOT REGISTER => use RegisterTaskJob(TaskJob)")
		}

		runner := Runner{
			taskJob:     taskType.TaskJobRef,
			taskStorage: taskService.TaskStorage,
		}
		fmt.Println("%v.taskType.CronFrequent:%v", taskType.Name, taskType.CronFrequent)
		errorSchedule := jobrunner.Schedule(taskType.CronFrequent, runner)
		if errorSchedule != nil {
			panic(errorSchedule)
		}
	}

}

func (taskService *TaskService) Load() {

	taskService.taskTypes = make(map[string]*TaskType)
	tasksTypes, error := taskService.TaskStorage.GetAllTaskType()
	if error != nil {
		panic(error)
	}
	for i := range tasksTypes {
		fmt.Println(tasksTypes[i].Name, tasksTypes[i].Description)
		taskService.taskTypes[tasksTypes[i].Name] = &tasksTypes[i]
	}

}
