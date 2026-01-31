package route

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/muh-hizbe/cashier-api/internal/handler"
)

func registerHealthRoutes(db *pgxpool.Pool) {
	hh := handler.NewHealthHandler(db)
	http.Handle("/health", hh)
}
