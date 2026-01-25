package route

import (
	"net/http"

	"github.com/muh-hizbe/cashier-api/internal/handler"
	"github.com/muh-hizbe/cashier-api/internal/repository"
)

func registerProductRoutes(w http.ResponseWriter, r *http.Request) {
	pr := repository.NewProductRepository()
	ph := handler.NewProductHandler(pr)
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
