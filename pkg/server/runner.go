package server

import (
	"encoding/json"
	"go.uber.org/zap"
	"time"
)

type Runner struct {
	jobTask     JobTask
	taskStorage TaskStorage
	logger      *zap.Logger
}

func (runner Runner) Run() {

	taskTypeName := runner.jobTask.GetTaskTypeName()
	runner.logger.Debug("START RUN", zap.String("TaskType", taskTypeName))
	tasks, error := runner.taskStorage.GetAllTaskNotStartedByType(taskTypeName)
	if error != nil {
		runner.logger.Error("ERROR", zap.Error(error))
		panic(error)
		return
	}
	for _, task := range tasks {
		runner.logger.Debug("EXECUTE RUN",
			zap.String("TaskType", taskTypeName),
			zap.String("UUID", task.Uuid))
		task.StartTime = time.Now()
		input, errorInput := runner.getInputMap(task)
		task.Error = errorInput
		runner.logger.Debug("ERROR INPUT JOB", zap.Error(errorInput),
			zap.String("INPUT", string(task.Input)))
		if task.Error == nil {
			output, errorExecute := runner.jobTask.Execute(input)
			if errorExecute != nil {
				task.Error = errorExecute
				runner.logger.Error("ERROR EXECUTE JOB", zap.Error(errorExecute))
			} else {
				task.Output, task.Error = runner.getOutputJson(task, output)
				runner.logger.Debug("ERROR OUTPUT JOB", zap.Error(task.Error))
			}
		}
		task.EndTime = time.Now()
		task.Finish = "S"
		runner.taskStorage.UpdateTask(task)
		var errorStr string = ""
		if task.Error != nil {
			errorStr = task.Error.Error()

		}
		runner.logger.Debug("END EXECUTE",
			zap.String("INPUT", string(task.Input)),
			zap.String("OUTPUT", string(task.Output)),
			zap.String("ERROR", errorStr),
			zap.String("UUID", task.Uuid))
	}

	runner.logger.Debug("FINISH RUN", zap.String("TaskType", taskTypeName))
}

func(runner Runner) getInputMap(task Task) (map[string]interface{}, error) {

	var input map[string]interface{}
	var errorResult error
	errorResult = json.Unmarshal(task.Input, &input)
	runner.logger.Debug("ERROR JSON INPUT:"  , zap.Error(errorResult))
	if errorResult == nil && task.TaskType.InputSchema != nil {
		errorResult = task.ValidateInput()
	}
	return input, errorResult
}

func(runner Runner) getOutputJson(task Task, output map[string]interface{}) ([]byte, error) {

	var outputByte []byte
	var errorResult error
	outputByte, errorResult = json.Marshal(output)
	runner.logger.Debug("ERROR OUTPUT JSON", zap.Error(errorResult))
	task.Output = outputByte
	if errorResult == nil && task.TaskType.OutputSchema != nil {
		errorResult = task.ValidateOutput()
	}
	return outputByte, errorResult
}
