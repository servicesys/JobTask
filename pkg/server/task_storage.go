package server

type TaskStorage interface {
	GetAllTaskType() ([]TaskType, error)
	GetAllTaskNotStartedByType(name string) ([]Task, error)
	SaveTask(task Task) error



}
