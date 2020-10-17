package main

// TaskService defines business operations on tasks.
type TaskService interface {
	ListTasks() []Task
	GetTask(ID string) Task
	CreateTask(task Task) Task
	UpdateTask(ID string, task Task) Task
	DeleteTask(ID string)
}

// NewTaskService creates a new task service implementation instance.
func NewTaskService() TaskService {
	return &taskService{}
}

// Holds any dependencies as shared state for the methods.
// Possible dependencies could be a database repository or external system interfaces.
// Concrete implementations of these interfaces are supplied in the main function.
type taskService struct{}

func (ts *taskService) ListTasks() []Task {
	return []Task{
		Task{
			ID:          "1abc1",
			Description: "Do chores",
			Status:      "opened",
		},
		Task{
			ID:          "2abc2",
			Description: "Clean the house",
			Status:      "closed",
		}}
}

func (ts *taskService) GetTask(ID string) Task {
	return Task{
		ID:          ID,
		Description: "Your task",
		Status:      "opened"}
}

func (ts *taskService) CreateTask(task Task) Task {
	return task
}

func (ts *taskService) UpdateTask(ID string, task Task) Task {
	return Task{
		ID:          ID,
		Description: task.Description,
		Status:      task.Status}
}

func (ts *taskService) DeleteTask(ID string) {
	return
}
