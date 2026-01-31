package route

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

func registerRoutes(db *pgxpool.Pool) {
	registerHealthRoutes(db)
	registerProductRoutes(db)
	registerCategoryRoutes(db)
}

func Init(db *pgxpool.Pool) {
	registerRoutes(db)
}
