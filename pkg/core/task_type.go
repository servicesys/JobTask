package core

type TaskType struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	InputSchema  []byte `json:"input_schema"`
	OutputSchema []byte `json:"output_schema"`
	CronFrequent string  `json:"cron_frequent"`
	TaskJobRef   TaskJob

}
