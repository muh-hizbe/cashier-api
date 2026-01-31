package route

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/muh-hizbe/cashier-api/internal/handler"
	"github.com/muh-hizbe/cashier-api/internal/repository"
	"github.com/muh-hizbe/cashier-api/internal/services"
)

func registerProductRoutes(db *pgxpool.Pool) {

	pr := repository.NewProductRepository(db)
	ps := services.NewProductService(pr)
	ph := handler.NewProductHandler(ps)
	http.Handle("/api/products", ph)
	http.Handle("/api/products/", ph)
	// http.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
	// 	switch r.Method {
	// 	case http.MethodGet:
	// 		ph.GetProducts(w, r)
	// 	case http.MethodPost:
	// 		ph.NewProduct(w, r)
	// 	default:
	// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	// 	}
	// })
	// http.HandleFunc("/api/products/", func(w http.ResponseWriter, r *http.Request) {
	// 	switch r.Method {
	// 	case http.MethodGet:
	// 		ph.GetProducts(w, r)
	// 	case http.MethodPost:
	// 		ph.NewProduct(w, r)
	// 	default:
	// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	// 	}
	// })
}
