package route

import (
	"net/http"

	"github.com/muh-hizbe/cashier-api/internal/handler"
)

func registerHealthRoutes(w http.ResponseWriter, r *http.Request) {
	hh := handler.NewHealthHandler()
	http.Handle("/health", hh)
}
