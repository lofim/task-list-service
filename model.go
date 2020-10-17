package main

// Task is a struct to hold task data
type Task struct {
	ID          string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status,omitempty"`
}
