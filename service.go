package main

import (
	"fmt"
	"math/rand"
	"time"
)

// TaskService defines business operations on tasks.
type TaskService interface {
	ListTasks() []Task
	GetTask(ID string) Task
	CreateTask(task Task) Task
	UpdateTask(ID string, task Task) Task
	DeleteTask(ID string)
}

func NewTaskService() TaskService {
	return &taskService{}
}

// Holds any dependencies as shared state for the methods.
// Possible dependencies could be a database repository or external system interfaces.
// Concrete implementations of these interfaces are supplied in the main function.
type taskService struct{}

func (ts *taskService) ListTasks() []Task {
	return []Task{
		{
			ID:          "1abc1",
			Description: "Do chores",
			Status:      "opened",
		},
		{
			ID:          "2abc2",
			Description: "Clean the house",
			Status:      "closed",
		}}
}

func (ts *taskService) GetTask(ID string) Task {
	return Task{
		ID:          ID,
		Description: "Your task",
		Status:      "opened",
	}
}

func (ts *taskService) CreateTask(task Task) Task {
	rand.Seed(time.Now().UnixNano())
	randomID := rand.Int()

	return Task{
		ID:          fmt.Sprint(randomID),
		Description: task.Description,
		Status:      "opened",
	}
}

func (ts *taskService) UpdateTask(ID string, task Task) Task {
	return Task{
		ID:          ID,
		Description: task.Description,
		Status:      task.Status,
	}
}

func (ts *taskService) DeleteTask(ID string) {}
