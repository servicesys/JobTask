package core

import "time"

type Task struct {
	Uuid	string	`json:"uuid"`
	TaskType TaskType	`json:"task_type"`
	Input	string	`json:"input"`
	Output	string	`json:"output"`
	History	string	`json:"history"`
	StartTime	time.Time	`json:"start_time"`
	EndTime	time.Time	`json:"end_time"`
	Error	string	`json:"error"`
	Finish	string	`json:"finish"`
	CreatedTime	time.Time	`json:"created_time"`
}
