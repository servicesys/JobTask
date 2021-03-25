package core

import (
	"encoding/json"
	"fmt"
	"time"
)

type Runner struct {
	taskJob     TaskJob
	taskStorage TaskStorage
}

func (runner Runner) Run() {

	taskTypeName := runner.taskJob.GetTaskTypeName()
	fmt.Println("Runner  + ..............", taskTypeName)
	fmt.Println(time.Now())

	tasks, error := runner.taskStorage.GetAllTaskNotStartedByType(taskTypeName)
	if error != nil {
		fmt.Println(error)
		//register error storage/
		return
	}

	for _, task := range tasks {
		fmt.Println("..............")
		fmt.Println(task.TaskType.Name)
		task.StartTime = time.Now()
		var input map[string]interface{}
		errJsonUnmarshal := json.Unmarshal(task.Input, &input)
		if errJsonUnmarshal != nil {
			task.Error = errJsonUnmarshal.Error()
			//register error history
		} else {
			output, errorExecute := runner.taskJob.Execute(input)
			if errorExecute != nil {
				task.Error = errorExecute.Error()
			} else {

				outputByte, errorJsonMarshal := json.Marshal(output)
				if errorJsonMarshal != nil {
					task.Error = errorJsonMarshal.Error()
				} else {
					task.Output = outputByte
				}
			}

		}

		task.EndTime = time.Now()
		task.Finish = "S"
		runner.taskStorage.UpdateTask(task)
		fmt.Println("..............")

	}
}
