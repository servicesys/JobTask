package server

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/qri-io/jsonschema"
	"strings"
	"time"
)

type Task struct {
	Uuid        string    `json:"uuid"`
	TaskType    TaskType  `json:"task_type"`
	Input       []byte    `json:"input"`
	Output      []byte    `json:"output"`
	History     []byte    `json:"history"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Error       error     `json:"error"`
	Finish      string    `json:"finish"`
	CreatedTime time.Time `json:"created_time"`
}

func (task Task) ValidateInput() error {

	valid, listErros := task.Validator(context.Background(), task.TaskType.InputSchema, task.Input)
	if !valid {
		return errors.New(strings.Join(listErros, "|"))
	}
	return nil
}

func (task Task) ValidateOutput() error {

	valid, listErros := task.Validator(context.Background(), task.TaskType.OutputSchema, task.Output)
	if !valid {
		return errors.New(listErros[0])
	}
	return nil
}

func (task Task) Validator(ctx context.Context, schemaBytes []byte, jsonBytes []byte) (bool, []string) {

	var listErrors = make([]string, 0)
	rs := &jsonschema.Schema{}

	if err := json.Unmarshal(schemaBytes, rs); err != nil {
		listErrors = append(listErrors, err.Error())
		return false, listErrors
	}

	errs, err := rs.ValidateBytes(ctx, jsonBytes)
	if err != nil {
		listErrors = append(listErrors, err.Error())
		return false, listErrors
	}

	if len(errs) > 0 {
		for _, e := range errs {
			listErrors = append(listErrors, e.Error())
		}
		return false, listErrors
	}

	return true, nil
}
