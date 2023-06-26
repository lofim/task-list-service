package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const UnexpectedErrorResponseJSON = "{\"statusCode\":500,\"errorMessage\":\"Unexpected error\"}"

// ControllerFunc defines an interface for all controller methods to be handled by genericHandlerWrapper
type ControllerFunc func(*http.Request) (interface{}, error)

// genericHandlerWrapper performs logging, error handling and response unification
func genericHandlerWrapper(handler ControllerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handlerResult, error := handler(r)

		if error != nil {
			// TODO: distinguish between different errors and set http status code appropriately
			// Example1: handle NotFound error and send 404 status code
			// Example2: handle Database connection errors and send 500 status code
			// Example3: handle validation errors and send 400 status code
			// Don't handle auth errors here. This should be done by the middleware
			log.Printf("Error executing generic response handler: %s", error.Error())
			respondError(w, error.Error(), http.StatusInternalServerError)

			return
		}

		// Don't write response body with StatusNoContent is used
		statusCode := getSuccessStatusCode(r.Method)
		if statusCode == http.StatusNoContent {
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(statusCode)

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
	errorResponse, err := json.Marshal(ErrorResponse{
		ErrorCode:    statusCode,
		ErrorMessage: error,
	})

	if err != nil {
		log.Printf("Error marshaling error response %v", err)
		errorResponse = []byte(UnexpectedErrorResponseJSON)
	}

	w.WriteHeader(statusCode)
	w.Write(errorResponse)
}

func getSuccessStatusCode(method string) int {
	var statusCode int
	switch method {

	case http.MethodGet:
		fallthrough
	case http.MethodPut:
		fallthrough
	case http.MethodPatch:
		statusCode = http.StatusOK
	case http.MethodPost:
		statusCode = http.StatusCreated
	case http.MethodDelete:
		statusCode = http.StatusNoContent
	default:
		statusCode = http.StatusNotImplemented
	}

	return statusCode
}
