package helper

import (
	"encoding/json"
	"net/http"
	"time"
)

type SuccessResponse struct {
	Timestamp string      `json:"timestamp"`
	Method    string      `json:"method"`
	Status    int         `json:"status"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

type ErrorResponse struct {
	Timestamp string `json:"timestamp"`
	Method    string `json:"method"`
	Status    int    `json:"status"`
	Message   string `json:"message"`
	Error     string `json:"error"`
}

func ResponseWriter(w http.ResponseWriter, r *http.Request, data any) {

	var status int
	message := "Request processed successfully"

	switch r.Method {
	case http.MethodGet:
		message = "Resource fetched successfully"
		status = http.StatusOK
	case http.MethodPost:
		message = "Resource created successfully"
		status = http.StatusCreated
	case http.MethodPut, http.MethodPatch:
		message = "Resource updated successfully"
		status = http.StatusOK
	case http.MethodDelete:
		message = "Resource deleted successfully"
		status = http.StatusOK
	default:
		status = http.StatusOK
	}

	response := SuccessResponse{
		Timestamp: time.Now().Format("2006-01-02T15:04:05"),
		Method:    r.Method,
		Status:    status,
		Message:   message,
		Data:      data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
func ErrorWriter(w http.ResponseWriter, r *http.Request, status int, err error) {

	statusMessageMap := map[int]string{
		// 4xx Client Errors
		http.StatusBadRequest:          "Invalid request parameters or malformed request",
		http.StatusUnauthorized:        "Authentication required or invalid credentials",
		http.StatusForbidden:           "You don't have permission to access this resource",
		http.StatusNotFound:            "The requested resource was not found",
		http.StatusMethodNotAllowed:    "HTTP method not allowed for this endpoint",
		http.StatusConflict:            "Resource conflict or duplicate entry",
		http.StatusUnprocessableEntity: "Request validation failed",
		http.StatusTooManyRequests:     "Too many requests, please try again later",

		// 5xx Server Errors
		http.StatusInternalServerError: "Internal server error",
		http.StatusNotImplemented:      "Feature not implemented",
		http.StatusServiceUnavailable:  "Service temporarily unavailable",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	message, ok := statusMessageMap[status]
	if !ok {
		message = "An unexpected error occurred. Please try again later or contact support."
	}

	response := ErrorResponse{
		Timestamp: time.Now().Format("2006-01-02T15:04:05"),
		Method:    r.Method,
		Status:    status,
		Message:   message,
		Error:     err.Error(),
	}

	json.NewEncoder(w).Encode(response)
}
