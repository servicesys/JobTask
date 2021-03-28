package infrastructure

import (
	"JobTask/pkg/server"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

type TaskStoragePostgres struct {
	dbConnection *pgx.Conn
}



func NewTaskClientStoragePostgres() server.TaskStorage {

	/**
	DB.HOST="localhost"
	DB.PORT="5432"
	DB.USER="valter"
	DB.PASS="valter"
	DB.NAME="app_sistema"
	*/
	connection := Connect("localhost", 5432, "valter", "valter", "app_sistema")
	storagePostgres := &TaskStoragePostgres{
		dbConnection: connection,
	}

	return storagePostgres
}
func Connect(host string, port int, user string, pass string, db string) *pgx.Conn {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, pass, db)

	dbConn, err := pgx.Connect(context.Background(), psqlInfo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		panic(err)
	}
	return dbConn

}



func doExecute(db *pgx.Conn, query string, args ...interface{}) error {

	_, err := db.Exec(context.Background(), query, args...)

	return err
}



func (t TaskStoragePostgres) GetAllTaskType() ([]server.TaskType, error) {
	panic("implement me")
}

func (t TaskStoragePostgres) GetAllTaskNotStartedByType(name string) ([]server.Task, error) {
	panic("implement me")
}

func (t TaskStoragePostgres) UpdateTask(task server.Task) error {
	panic("implement me")
}