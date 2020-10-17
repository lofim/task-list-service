package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// ControllerFunc defines an interface for all controller methods to be handled by generticHandlerWrapper
type ControllerFunc func(*http.Request) (interface{}, error)

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "up")
}

// genericHandlerWrapper performs logging, error handling and response unification
func genericHandlerWrapper(handler ControllerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Invoked api: %s %s", r.Method, r.RequestURI)

		handlerResult, error := handler(r)
		if error != nil {
			// TODO: distinguish between different errors and set http status code appropriately
			// Example1: handle NotFound error and send 404 status code
			// Example2: handle Database connection erros and send 500 status code
			// Example3: handle validation errors and send 400 status code
			// Don't handle auth errors here. This should be done by the middleware
			log.Printf("Error executing generic response handler: %s", error.Error())
			respondError(w, error.Error(), http.StatusInternalServerError)

			return
		}

		defer log.Printf("Api: %s %s, response: %v", r.Method, r.RequestURI, handlerResult)

		statusCode := getStatusCode(r.Method)
		w.WriteHeader(statusCode)

		// Don't write response body with StatusNoContent is use
		if statusCode == http.StatusNoContent {
			return
		}

		jsonEncoder := json.NewEncoder(w)
		error = jsonEncoder.Encode(handlerResult)
		if error != nil {
			log.Printf("Error marshaling response body to json: %s", error.Error())
			respondError(w, error.Error(), http.StatusInternalServerError)

			return
		}
	}
}

func respondError(w http.ResponseWriter, error string, statusCode int) {
	errorResponse, _ := json.Marshal(ErrorResponse{
		ErrorCode:    http.StatusBadRequest,
		ErrorMessage: error,
	})

	w.WriteHeader(statusCode)
	w.Write(errorResponse)
}

func getStatusCode(method string) int {
	var statusCode int
	switch method {
	case http.MethodPost:
		statusCode = http.StatusCreated
	case http.MethodGet:
		statusCode = http.StatusOK
	case http.MethodDelete:
		statusCode = http.StatusNoContent
	case http.MethodPut:
		statusCode = http.StatusOK
	default:
		statusCode = http.StatusNotImplemented
	}

	return statusCode
}
