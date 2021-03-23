package core

import (
	"fmt"
	"time"
)

type Runner struct {
	taskJob TaskJob
	taskStorage TaskStorage
}

func (runner Runner) Run() {


	taskTypeName := runner.taskJob.GetTaskTypeName()
	/*
	tasks , error :=runner.taskStorage.GetAllTaskNotStartedByType(taskTypeName)
	if error !=nil {

		fmt.Println(error)
		//register error storage/
        return
	}

	for  _, task:= range  tasks {

		fmt.Println(task)
		task.StartTime = time.Now()
		//input := task.Input
		input := make(map[string]interface{})
		input["RUN"] = "RUNNER"
		output, error := runner.taskJob.Execute(input)
		fmt.Println(error)
		fmt.Println(time.Now())
		fmt.Println(output)

		if error!=nil{
			task.Error = error.Error()
		}else {
			task.Output = output
		}
		task.EndTime = time.Now()
		task.Finish = "S"
		runner.taskStorage.Update(task)
	}
*/
	fmt.Println("Runner  + .............." , taskTypeName)
	//Execute(input map[string]interface{}) (map[string]interface{}, error)
	input := make(map[string]interface{})
	input["RUN"] = "RUNNER"
	output, error := runner.taskJob.Execute(input)
	fmt.Println(error)
	fmt.Println(time.Now())
	fmt.Println(output)
	fmt.Println("..............")
}


