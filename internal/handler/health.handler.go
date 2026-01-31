package handler

import (
	"net/http"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/muh-hizbe/cashier-api/internal/response"
)

type HealthHandler struct {
	DB *pgxpool.Pool
}

func NewHealthHandler(db *pgxpool.Pool) *HealthHandler {
	return &HealthHandler{
		DB: db,
	}
}

func (h *HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	path := strings.TrimPrefix(r.URL.Path, "/health")

	switch {
	case path == "" || path == "/":
		h.CheckHealth(w, r, h.DB)
	default:
		response.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *HealthHandler) CheckHealth(w http.ResponseWriter, r *http.Request, db *pgxpool.Pool) {
	ctx := r.Context()

	result := map[string]any{
		"app":      "ok",
		"database": "ok",
	}

	if err := db.Ping(ctx); err != nil {
		result["database"] = "failed"
	}

	response.Success(w, "APP running", result)
}
