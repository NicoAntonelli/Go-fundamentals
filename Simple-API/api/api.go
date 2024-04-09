package api

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Code    int
	Message string
}

// Error responses for HTTP requests
func writeError(w http.ResponseWriter, message string, code int) {
	// Set error struct
	resp := ErrorResponse{
		Code:    code,
		Message: message,
	}

	// Set headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	// Write response (JSON)
	json.NewEncoder(w).Encode(resp)
}

// Wrappers for writeError
var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "Internal server error", http.StatusInternalServerError)
	}
)
