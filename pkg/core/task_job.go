package core

type TaskJob interface {

	Execute(input map[string]interface{}) (map[string]interface{}, error)
	GetTaskTypeName() string
}
