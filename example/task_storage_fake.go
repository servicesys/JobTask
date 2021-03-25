package example

import (
	"JobTask/pkg/core"
	"errors"
	"fmt"
	"time"
)

type TaskStorageFake struct {
}

func (t TaskStorageFake) GetAllTaskType() ([]core.TaskType, error) {

	helloTaskType := core.TaskType{
		Name:         "HELLO",
		Description:  "Hello for users",
		InputSchema:  nil,
		OutputSchema: nil,
		CronFrequent: "@every 5s",
	}

	worldTaskType := core.TaskType{
		Name:         "WORLD",
		Description:  "World for users",
		InputSchema:  nil,
		OutputSchema: nil,
		CronFrequent: "@every 25s",
	}

	listTaskTypes := []core.TaskType{helloTaskType, worldTaskType}

	return listTaskTypes, nil

}

func (t TaskStorageFake) GetAllTaskNotStartedByType(name string) ([]core.Task, error) {

	if name == "HELLO" {
		return getHelloTasks(), nil
	}

	if name == "WORLD" {
		return getWorldTasks(), nil
	}

	return nil, errors.New("FAIL " + name)
}

func (t TaskStorageFake) UpdateTask(task core.Task) error {

	fmt.Println("//UPDATE:" + task.TaskType.Name)
	fmt.Println(task.StartTime, task.EndTime, task.Finish)
	fmt.Println(string(task.Output))
	fmt.Println(task.Error)
	fmt.Println("//UPDATE:" + task.TaskType.Name)

	return nil

}

func getHelloTasks() []core.Task {

	textoJSon := `
 { "titulo" : " Hello world task job input" , "texto" :  "HELLO"
		 } ;`

	input := []byte(textoJSon)

	listHelloTask := make([]core.Task, 5)

	for i := 0; i < 5; i++ {

		listHelloTask = append(listHelloTask, core.Task{
			Uuid: "hello" + string(i),
			TaskType: core.TaskType{
				Name:         "HELLO",
				Description:  "",
				InputSchema:  nil,
				OutputSchema: nil,
				CronFrequent: "",
				TaskJobRef:   nil,
			},
			Input:       input,
			Output:      nil,
			History:     nil,
			StartTime:   time.Time{},
			EndTime:     time.Time{},
			Error:       "",
			Finish:      "",
			CreatedTime: time.Time{},
		})

	}

	return listHelloTask

}

func getWorldTasks() []core.Task {

	textoJSon := `
 { "titulo" : " World task job input" , "texto" :  "WORLD"
		 } `

	input := []byte(textoJSon)

	listHelloTask := make([]core.Task, 5)

	for i := 0; i < 2; i++ {

		listHelloTask = append(listHelloTask, core.Task{
			Uuid: "world" + string(i),
			TaskType: core.TaskType{
				Name:         "WORLD",
				Description:  "",
				InputSchema:  nil,
				OutputSchema: nil,
				CronFrequent: "",
				TaskJobRef:   nil,
			},
			Input:       input,
			Output:      nil,
			History:     nil,
			StartTime:   time.Time{},
			EndTime:     time.Time{},
			Error:       "",
			Finish:      "",
			CreatedTime: time.Time{},
		})

	}

	return listHelloTask

}
