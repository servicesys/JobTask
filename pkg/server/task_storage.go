package server

import "context"

type TaskStorage interface {
	GetAllTaskType() ([]TaskType, error)
	GetAllTaskNotStartedByType(ctx context.Context , name string) ([]Task, error)
	SaveTask(task Task) error



}
