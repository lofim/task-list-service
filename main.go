package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Main function does most setup here.
	// Since we don't use IoC container we need to create all dependencies manually
	// taskService is a dependency of taskController
	// taskService holds the business logic and is implementation agnostic, decoupled using an interface
	// taskController is a collection of functions/API handlers, we could also apply interface abstraction here
	// taskController does not implement any business logic only request/response mapping

	// Sample of creating an object instance via constructor function
	taskService := NewTaskService()

	// Sample of creating an object instance via struct literal
	// Also shows inversion of control (IoC) of taskService dependency
	taskController := TaskController{taskService}

	// Initialize router for version 1 of the API
	apiV1 := chi.NewRouter()

	// register controller handlers wrapped in genericHandlerWrapper for error handling
	apiV1.Get("/tasks", genericHandlerWrapper(taskController.listTasks))
	apiV1.Post("/tasks", genericHandlerWrapper(taskController.createTask))
	apiV1.Delete("/tasks/{id}", genericHandlerWrapper(taskController.deleteTask))
	apiV1.Get("/tasks/{id}", genericHandlerWrapper(taskController.getTask))
	apiV1.Put("/tasks/{id}", genericHandlerWrapper(taskController.updateTask))

	r := chi.NewRouter()
	r.Use(middleware.Heartbeat("/health"))
	r.Use(middleware.Logger)
	r.Mount("/api/v1", apiV1)

	log.Println("Starting server on port 8080 ...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
