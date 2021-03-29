package main

import (

	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/servicesys/JobTask/example"
	"github.com/servicesys/JobTask/pkg/infrastructure"
	"github.com/servicesys/JobTask/pkg/server"
	"os"
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
	connection := Connect("localhost", 5432, "valter", "valter", "app_sistema")
	taskStoragePostgres := infrastructure.NewTaskClientStoragePostgres(connection)
	taskService := server.NewTaskService(taskStoragePostgres) //
	//taskService := server.NewTaskService(example.TaskStorageFake{})
	taskService.RegisterTaskJob(example.TaskJobHello{})
	taskService.RegisterTaskJob(example.TaskJobWorld{})
	taskService.Start()

	for {
		time.Sleep(time.Second * 180)
	}
}

func Connect(host string, port int, user string, pass string, db string) *pgxpool.Pool {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, pass, db)

	dbConn, err:=pgxpool.Connect(context.Background(), psqlInfo)
	//dbConn, err := pgx.Connect(context.Background(), psqlInfo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		panic(err)
	}
	return dbConn

}
