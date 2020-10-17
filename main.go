package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Main function does most setup here.
	// Since we don't use Ioc container we need to create all dependencies manually
	// taskService is a dependency of taskController
	// taskService holds the bussiness logic and is implementation agnostic, decoupled using an interface
	// taskController is a collection of functions/API handlers, we could also apply interface abstraction here
	// taskController does not implement any business logic only request/response mapping

	// Sample of creating an object instance via contructor function
	taskService := NewTaskService()

	// Sample of createing an object instance via struct literal
	// Also shows inversion of controll (IoC) of taskService dependency
	taskController := TaskController{taskService}

	// Initialize router for version 1 of the API
	r := mux.NewRouter()
	apiV1 := r.PathPrefix("/api/v1").Subrouter()

	// register controller handlers wrapped in genericHandlerWrapper for error handling
	apiV1.HandleFunc("/tasks", genericHandlerWrapper(taskController.listTasks)).Methods(http.MethodGet)
	apiV1.HandleFunc("/tasks", genericHandlerWrapper(taskController.createTask)).Methods(http.MethodPost)
	apiV1.HandleFunc("/tasks/{id}", genericHandlerWrapper(taskController.deleteTask)).Methods(http.MethodDelete)
	apiV1.HandleFunc("/tasks/{id}", genericHandlerWrapper(taskController.getTask)).Methods(http.MethodGet)
	apiV1.HandleFunc("/tasks/{id}", genericHandlerWrapper(taskController.updateTask)).Methods(http.MethodPut)

	// register utility handlers such as healh check
	r.HandleFunc("/health", health).Methods(http.MethodGet)
	http.Handle("/", r)

	log.Println("Starting server on port 8080 ...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
