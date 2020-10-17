package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "up")
}

// genericHandlerWrapper performs logging, error handling and response unification
func genericHandlerWrapper(handler func(*http.Request) (interface{}, error)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Invoked api: %s %s", r.Method, r.RequestURI)

		handlerResult, error := handler(r)
		if error != nil {
			// TODO: distinguish between different errors and set http status code appropriately
			log.Printf("Error executing generic response handler: %s", error.Error())
			respondError(w, error.Error(), http.StatusInternalServerError)

			return
		}

		jsonEncoder := json.NewEncoder(w)

		w.WriteHeader(getStatusCode(r.Method))
		error = jsonEncoder.Encode(handlerResult)
		if error != nil {
			log.Printf("Error marshaling response body to json: %s", error.Error())
			respondError(w, error.Error(), http.StatusInternalServerError)

			return
		}

		log.Printf("Api: %s %s, response: %v", r.Method, r.RequestURI, handlerResult)
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
		statusCode = http.StatusInternalServerError
	}

	return statusCode
}
