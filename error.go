package main

// ErrorResponse is the data trasfer object for errors
type ErrorResponse struct {
	ErrorCode    int    `json:"statusCode"`
	ErrorMessage string `json:"errorMessage"`
}
