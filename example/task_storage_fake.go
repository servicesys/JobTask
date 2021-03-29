package example

import (
	"JobTask/pkg/server"
	"errors"
	"fmt"
	"time"
)

type TaskStorageFake struct {
}

func (t TaskStorageFake) GetAllTaskType() ([]server.TaskType, error) {

	helloTaskType := server.TaskType{
		Name:         "HELLO",
		Description:  "Hello for users",
		InputSchema:  nil,
		OutputSchema: nil,
		CronFrequent: "@every 5s",
	}

	worldTaskType := server.TaskType{
		Name:         "WORLD",
		Description:  "World for users",
		InputSchema:  nil,
		OutputSchema: nil,
		CronFrequent: "@every 15s",
	}

	listTaskTypes := []server.TaskType{helloTaskType, worldTaskType} //helloTaskType,

	return listTaskTypes, nil

}

func (t TaskStorageFake) GetAllTaskNotStartedByType(name string) ([]server.Task, error) {

	if name == "HELLO" {
		return getHelloTasks(), nil
	}

	if name == "WORLD" {
		return getWorldTasks(), nil
	}

	return nil, errors.New("FAIL " + name)
}

func (t TaskStorageFake) SaveTask(task server.Task) error {

	fmt.Println("//UPDATE:" + task.TaskType.Name)
	fmt.Println(task.StartTime, task.EndTime, task.Finish)
	fmt.Println("INPUT:")
	fmt.Println(string(task.Input))
	fmt.Println("OUTPUT:")
	fmt.Println(string(task.Output))
	fmt.Println("ERROR:")
	fmt.Println(task.Error)
	fmt.Println("//UPDATE:" + task.TaskType.Name)

	return nil

}

func getHelloTasks() []server.Task {

	textoJSon := ` { "title" : " Hello world task job input" , "text" :  "HELLO"} `

	input := []byte(textoJSon)

	listHelloTask := make([]server.Task, 2)

	for i := 0; i < 2; i++ {

		listHelloTask[i] = server.Task{
			Uuid: "hello" + string(i),
			TaskType: server.TaskType{
				Name:         "HELLO",
				Description:  "",
				InputSchema:  nil,
				OutputSchema: nil,
				CronFrequent: "",
				JobTaskRef:   nil,
			},
			Input:       input,
			Output:      nil,
			History:     nil,
			StartTime:   time.Time{},
			EndTime:     time.Time{},
			Error:       nil,
			Finish:      "",
			CreatedTime: time.Time{},
		}

	}
	fmt.Println(len(listHelloTask))
	return listHelloTask

}

func getSchema() []byte {

	SCHEMA := `{
     "$id": "https://qri.io/schema/",
    "$comment" : "sample comment",
    "title": "Texto Blog",
    "type": "object",
	"properties": {
		"title": {

			"title": "Titulo",
			"type": "string",
			"default": "",
			"examples": [
				"Este e um texto de exemplo"
			],
			"pattern": "^.*$"
		},
		"text": {

			"title": "Texto",
			"type": "string",
			"default": "",
			"examples": [
				"<p>Este e  o corpot do texto texto de exemplo</p>"
			],
			"pattern": "^.*$"
		}
	},
	"required": [
		"title",
		"text"
	]
}
`
	return []byte(SCHEMA)
}

func getWorldTasks() []server.Task {

	textoJSon := `{ "title" : " World task job input" , "text" :  "WORLD"}`

	input := []byte(textoJSon)

	listHelloTask := make([]server.Task, 2)

	for i := 0; i < 2; i++ {

		listHelloTask[i] = server.Task{
			Uuid: "world" + string(i),
			TaskType: server.TaskType{
				Name:         "WORLD",
				Description:  "",
				InputSchema:  getSchema(),
				OutputSchema: getSchema(),
				CronFrequent: "",
				JobTaskRef:   nil,
			},
			Input:       input,
			Output:      nil,
			History:     nil,
			StartTime:   time.Time{},
			EndTime:     time.Time{},
			Error:       nil,
			Finish:      "",
			CreatedTime: time.Time{},
		}

	}

	return listHelloTask

}
