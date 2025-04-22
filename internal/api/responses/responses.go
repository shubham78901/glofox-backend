// File: internal/api/responses/responses.go

package responses

import (
	"encoding/json"
	"net/http"
)

// Response represents a standard API response
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Count   int         `json:"count,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// WriteJSON writes a JSON response to the http.ResponseWriter
func WriteJSON(w http.ResponseWriter, statusCode int, response Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

// SuccessResponse sends a success response
func SuccessResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	WriteJSON(w, statusCode, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// CreatedResponse sends a 201 Created response
func CreatedResponse(w http.ResponseWriter, message string, data interface{}) {
	SuccessResponse(w, http.StatusCreated, message, data)
}

// OKResponse sends a 200 OK response
func OKResponse(w http.ResponseWriter, data interface{}) {
	WriteJSON(w, http.StatusOK, Response{
		Success: true,
		Data:    data,
	})
}

// ListResponse sends a 200 OK response with a count field
func ListResponse(w http.ResponseWriter, data interface{}, count int) {
	WriteJSON(w, http.StatusOK, Response{
		Success: true,
		Count:   count,
		Data:    data,
	})
}

// ErrorResponse sends an error response
func ErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	WriteJSON(w, statusCode, Response{
		Success: false,
		Message: message,
	})
}

// BadRequestResponse sends a 400 Bad Request response
func BadRequestResponse(w http.ResponseWriter, message string) {
	ErrorResponse(w, http.StatusBadRequest, message)
}

// NotFoundResponse sends a 404 Not Found response
func NotFoundResponse(w http.ResponseWriter, message string) {
	ErrorResponse(w, http.StatusNotFound, message)
}

// InternalServerErrorResponse sends a 500 Internal Server Error response
func InternalServerErrorResponse(w http.ResponseWriter) {
	ErrorResponse(w, http.StatusInternalServerError, "Internal server error")
}
