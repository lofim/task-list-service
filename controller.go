package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// TaskController holds the API request/response mappings
type TaskController struct {
	// additional dependencies
	// logger, common mapping functionality, etc.
	taskService TaskService
}

func (tc *TaskController) listTasks(r *http.Request) (interface{}, error) {
	// call service layer
	// ignore service layer errors for now
	return tc.taskService.ListTasks(), nil
}

func (tc *TaskController) createTask(r *http.Request) (interface{}, error) {
	// map input data to service layer data (unmarshal json to model)
	jsonDecoder := json.NewDecoder(r.Body)
	jsonDecoder.DisallowUnknownFields()

	var task Task
	error := jsonDecoder.Decode(&task)
	if error != nil {
		return Task{}, fmt.Errorf("error unmarshaling json body: %w", error)
	}

	// very naive validation
	// possibly use a validation library like this one
	// https://github.com/go-playground/validator
	if task.Description == "" {
		return Task{}, errors.New("Task body is missing description")
	}

	// call service layer
	// ignoring service layer errors for now
	return tc.taskService.CreateTask(task), nil
}

func (tc *TaskController) getTask(r *http.Request) (interface{}, error) {
	taskID := chi.URLParam(r, "id")
	return tc.taskService.GetTask(taskID), nil
}

func (tc *TaskController) updateTask(r *http.Request) (interface{}, error) {
	taskID := chi.URLParam(r, "id")
	jsonDecoder := json.NewDecoder(r.Body)
	jsonDecoder.DisallowUnknownFields()

	var task Task
	error := jsonDecoder.Decode(&task)
	if error != nil {
		return nil, fmt.Errorf("error unmarshaling json body: %w", error)
	}

	return tc.taskService.UpdateTask(taskID, task), nil
}

func (tc *TaskController) deleteTask(r *http.Request) (interface{}, error) {
	taskID := chi.URLParam(r, "id")
	tc.taskService.DeleteTask(taskID)

	return nil, nil
}
