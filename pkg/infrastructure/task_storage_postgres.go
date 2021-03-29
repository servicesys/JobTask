package infrastructure

import (

	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/servicesys/JobTask/pkg/server"
	"strings"
)

type TaskStoragePostgres struct {
	dbConnection *pgxpool.Pool
}

func NewTaskClientStoragePostgres(connection *pgxpool.Pool) server.TaskStorage {

	storagePostgres := &TaskStoragePostgres{
		dbConnection: connection,
	}

	return storagePostgres
}

func doExecute(db *pgxpool.Pool, query string, args ...interface{}) error {

	_, err := db.Exec(context.Background(), query, args...)

	return err
}

func (t TaskStoragePostgres) GetAllTaskType() ([]server.TaskType, error) {
	strQuery :=
		`SELECT name , description, input_schema, output_schema, 
         cron_frequent, enable FROM job_task.task_type WHERE enable='S' OR  enable IS NULL;`

	taskTypes := make([]server.TaskType, 0)
	rows, errQuery := t.dbConnection.Query(context.Background(), strQuery)
	taskType := server.TaskType{}

	if errQuery != nil {
		return taskTypes, errQuery
	}
	for rows.Next() {
		taskType = server.TaskType{}
		strName := sql.NullString{}
		enable := sql.NullString{}
		errorScan := rows.Scan(&strName,
			&taskType.Description,
			&taskType.InputSchema,
			&taskType.OutputSchema,
			&taskType.CronFrequent,
			&enable)

		if errorScan != nil {
			return taskTypes, errorScan
		}
		taskType.Enable = enable.String == "S"
		taskType.Name = strings.TrimSpace(strName.String)
		taskTypes = append(taskTypes, taskType)
	}
	defer rows.Close()
	return taskTypes, nil
}

func (t TaskStoragePostgres) GetAllTaskNotStartedByType(ctx context.Context, name string) ([]server.Task, error) {

	strQuery :=
		`SELECT uuid, task_type_name , input , output , start_time, end_time, error, finish, created_time,
         tt.description, tt.input_schema, tt.output_schema, tt.cron_frequent
         FROM job_task.task t 
         INNER JOIN job_task.task_type tt ON (t.task_type_name=tt.name)
         WHERE (finish is null OR finish ='N') AND (enable='S' OR enable IS NULL) AND task_type_name = $1;`

	//fmt.Println(strQuery)
	var tasks []server.Task
	rows, errQuery := t.dbConnection.Query(ctx, strQuery, name)
	if errQuery != nil {
		panic(errQuery)
	}
	for rows.Next() {
		task := server.Task{}
		task.TaskType = server.TaskType{}
		startTime := sql.NullTime{}
		endTime := sql.NullTime{}
		strError := sql.NullString{}
		strFinish := sql.NullString{}
		strName := sql.NullString{}
		errorScan := rows.Scan(&task.Uuid,
			&strName,
			&task.Input,
			&task.Output,
			&startTime,
			&endTime,
			&strError,
			&strFinish,
			&task.CreatedTime,
			&task.TaskType.Description,
			&task.TaskType.InputSchema,
			&task.TaskType.OutputSchema,
			&task.TaskType.CronFrequent)

		if errorScan != nil {
			return tasks, errorScan
		}
		task.TaskType.Name = strings.TrimSpace(strName.String)
		task.StartTime = startTime.Time
		task.EndTime = endTime.Time
		task.Error = errors.New(strError.String)
		task.Finish = strFinish.String
		tasks = append(tasks, task)
	}
	rows.Close()
	return tasks, nil
}

func (t TaskStoragePostgres) SaveTask(task server.Task) error {

	query := `UPDATE job_task.task 
              SET output=$1 , history= $2, start_time=$3, end_time= $4 , error=$5, finish=$6 WHERE uuid=$7;`

	var strError string
	if task.Error != nil {
		strError = task.Error.Error()
	}
	err := doExecute(t.dbConnection, query, task.Output, task.History, task.StartTime, task.EndTime,
		strError, task.Finish, task.Uuid)
	return err
}
