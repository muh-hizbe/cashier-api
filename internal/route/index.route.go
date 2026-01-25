package route

import "net/http"

func registerRoutes(w http.ResponseWriter, r *http.Request) {
	registerHealthRoutes(w, r)
	registerProductRoutes(w, r)
	registerCategoryRoutes(w, r)
}

func Init(w http.ResponseWriter, r *http.Request) {
	registerRoutes(w, r)
}
