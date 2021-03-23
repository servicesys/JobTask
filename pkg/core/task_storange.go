package core


type TaskStorage interface {


	GetAllTaskNotStartedByType(name string) ([]Task, error)
	UpdateTask(task Task) error

	/*CreateTaskType(taskType TaskType) error
	UpdateTaskType(taskType TaskType) error
	DeleteTaskType(taskType TaskType) error
	GetTaskTypeByName(name string) (TaskType, error)

	AddTask(task Task) error
	UpdateTask(task Task) error
	DeleteTask(task Task) error*/

}
