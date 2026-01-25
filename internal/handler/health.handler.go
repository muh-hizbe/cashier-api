package handler

import (
	"net/http"
	"strings"

	"github.com/muh-hizbe/cashier-api/internal/response"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	path := strings.TrimPrefix(r.URL.Path, "/health")

	switch {
	case path == "" || path == "/":
		h.CheckHealth(w, r)
	default:
		response.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *HealthHandler) CheckHealth(w http.ResponseWriter, r *http.Request) {
	response.Success(w, "API Running", nil)
}
