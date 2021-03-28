package server

type JobTask interface {
	Execute(input map[string]interface{}) (map[string]interface{}, error)
	GetTaskTypeName() string
}
