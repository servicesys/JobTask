package example

import (
	"fmt"
	"time"
)

type TaskJobWorld struct {
}

func (taskJobWorld TaskJobWorld) Execute(input map[string]interface{}) (map[string]interface{}, error) {

	output := make(map[string]interface{})
	output["OLA"] = "WORLD " + fmt.Sprintf(" %v ", input["titulo"])
	fmt.Println("TASK WORLD..............")
	fmt.Println(time.Now())
	fmt.Println("..............")
	return output, nil
}


func (taskJobWorld TaskJobWorld) GetTaskTypeName() string {

	return "WORLD"
}