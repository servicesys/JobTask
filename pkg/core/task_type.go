package core


type TaskType struct {
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	InputSchema	string	`json:"input_schema"`
	OutputSchema	string	`json:"output_schema"`

}
