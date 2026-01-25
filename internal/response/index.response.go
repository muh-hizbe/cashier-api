package response

import (
	"encoding/json"
	"net/http"
)

func NewResponse(statusCode int, data any, message string) map[string]any {
	return map[string]any{
		"status":  statusCode,
		"data":    data,
		"message": message,
	}
}

func Error(w http.ResponseWriter, msg string, status ...int) {
	if len(status) == 0 {
		status = []int{http.StatusInternalServerError}
	}
	w.WriteHeader(status[0])
	json.NewEncoder(w).Encode(NewResponse(status[0], nil, msg))
}

func Success(w http.ResponseWriter, msg string, data any, status ...int) {
	if len(status) == 0 {
		status = []int{http.StatusOK}
	}
	w.WriteHeader(status[0])
	json.NewEncoder(w).Encode(NewResponse(status[0], data, msg))
}
