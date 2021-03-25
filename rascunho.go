package main

import (
	"JobTask/example"
	"JobTask/pkg/core"
	"time"
)

func main() {

	/*
	textoJSon := `
 { "titulo" : "Este e um texto de exemplo -NOVO - ola" , "texto" :  "Corpo do texto texto de exemplo"
		 }`

	data := []byte(textoJSon)
	var p map[string]interface{}
	err := json.Unmarshal(data, &p)
	fmt.Println(err)
	fmt.Println(p)
	fmt.Println("TESTANDO")
    */

	taskService := core.TaskService{TaskStorage : example.TaskStorageFake{}}
	taskService.Load()
	taskService.RegisterTaskJob(example.TaskJobHello{})
	taskService.RegisterTaskJob(example.TaskJobWorld{})
	taskService.Start()


	for {
		time.Sleep(time.Second * 180)
	}
}
