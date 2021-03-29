package example

import (
	"fmt"
	"time"
)

type TaskJobHello struct {
}

func (taskJobHello TaskJobHello) Execute(input map[string]interface{}) (map[string]interface{}, error) {

	output := make(map[string]interface{})
	output["HELLO"] = " ---" + fmt.Sprintf(" %v ", input["title"])
	fmt.Println("TASK HELLO..............")
	fmt.Println(time.Now())
	fmt.Println("..............")
	return output, nil //errors.New("PROCESSAMENTO INCORRETO:" + time.Now().String() )
}

func (taskJobHello TaskJobHello) GetTaskTypeName() string {

	return "HELLO"
}
