package server

import (
	"encoding/json"
	"github.com/bamzi/jobrunner"
	"go.uber.org/zap"
	"log"
)

type TaskService struct {
	TaskStorage TaskStorage
	taskTypes   map[string]*TaskType
	logger      *zap.Logger
}

func NewTaskService(taskStorage TaskStorage) *TaskService {

	rawJSON := []byte(`{
	  "level": "debug",
	  "encoding": "json",
	  "outputPaths": ["stdout", "/tmp/logs"],
	  "errorOutputPaths": ["stderr"],
	  "initialFields": {"JOBTASK": "V1.0"},
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase"
	  }
	}`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger, errorLog := cfg.Build()
	if errorLog != nil {
		log.Fatalf("can't initialize zap logger: %v", errorLog)
	}
	defer logger.Sync()
	taskService := TaskService{TaskStorage: taskStorage, logger: logger}
	taskService.load()
	return &taskService

}
func (taskService *TaskService) RegisterTaskJob(jobTask JobTask) {

	keyTaskJob := jobTask.GetTaskTypeName()
	taskService.logger.Info("TASK TYPE:[" + keyTaskJob)
	if taskType, found := taskService.taskTypes[keyTaskJob]; found {
		taskType.JobTaskRef = jobTask
	} else {
		panic("NOT EXIST TASK TYPE WITH NAME: [" + keyTaskJob + "]")
	}
}

func (taskService *TaskService) Start() {

	taskService.logger.Info("TaskService:Start")
	jobrunner.Start(2)

	for _, taskType := range taskService.taskTypes {

		if taskType.JobTaskRef == nil {
			taskService.logger.Panic("TaskJob : [" + taskType.Name + "] NOT REGISTER => use RegisterTaskJob(TaskJob)")
			panic("TaskJob : [" + taskType.Name + "] NOT REGISTER => use RegisterTaskJob(TaskJob)")

		}

		runner := Runner{
			jobTask:     taskType.JobTaskRef,
			taskStorage: taskService.TaskStorage,
			logger:      taskService.logger,
		}
		taskService.logger.Info("Runner:["+taskType.Name+"]",
			zap.String("CronFrequent", taskType.CronFrequent),
			zap.String("Name", taskType.Name))
		errorSchedule := jobrunner.Schedule(taskType.CronFrequent, runner)
		if errorSchedule != nil {
			panic(errorSchedule)
		}
	}

}

func (taskService *TaskService) load() {

	taskService.taskTypes = make(map[string]*TaskType)
	tasksTypes, error := taskService.TaskStorage.GetAllTaskType()
	if error != nil {
		panic(error)
	}
	for i := range tasksTypes {
		taskService.logger.Info(tasksTypes[i].Name,
			zap.String("CronFrequent", tasksTypes[i].CronFrequent),
			zap.String("Name", tasksTypes[i].Name))
		taskService.taskTypes[tasksTypes[i].Name] = &tasksTypes[i]
	}

}
