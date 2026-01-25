package route

import (
	"net/http"

	"github.com/muh-hizbe/cashier-api/internal/handler"
	"github.com/muh-hizbe/cashier-api/internal/repository"
)

func registerCategoryRoutes(w http.ResponseWriter, r *http.Request) {
	cr := repository.NewCategoryRepository()
	ch := handler.NewCategoryHandler(cr)
	http.Handle("/api/categories", ch)
	http.Handle("/api/categories/", ch)
}
