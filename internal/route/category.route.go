package route

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/muh-hizbe/cashier-api/internal/handler"
	"github.com/muh-hizbe/cashier-api/internal/repository"
	"github.com/muh-hizbe/cashier-api/internal/services"
)

func registerCategoryRoutes(db *pgxpool.Pool) {
	cr := repository.NewCategoryRepository(db)
	cs := services.NewCategoryService(cr)
	ch := handler.NewCategoryHandler(cs)
	http.Handle("/api/categories", ch)
	http.Handle("/api/categories/", ch)
}
