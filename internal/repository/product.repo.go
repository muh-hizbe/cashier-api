package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/muh-hizbe/cashier-api/db/queries"
	"github.com/muh-hizbe/cashier-api/internal/domain"
	"github.com/muh-hizbe/cashier-api/internal/model"
)

type ProductRepository struct {
	DB *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

func (r *ProductRepository) GetProducts() ([]model.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	rows, err := r.DB.Query(ctx, queries.GET_PRODUCTS_WITH_CATEGORY)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	products, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[model.Product])
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) GetProduct(id int) (*model.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	rows, err := r.DB.Query(ctx, queries.GET_PRODUCT_BY_ID_WITH_CATEGORY, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	product, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[model.Product])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		log.Println(err)
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) CreateProduct(payload *model.Product) (*model.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	rows, err := r.DB.Query(ctx, queries.INSERT_PRODUCT, payload.Name, payload.Price, payload.Stock, payload.CategoryID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	product, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[model.Product])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		log.Println(err)
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) UpdateProduct(payload *model.Product, id int) (*model.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	rows, err := r.DB.Query(ctx, queries.UPDATE_PRODUCT, id, payload.Name, payload.Price, payload.Stock, *payload.CategoryID)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23503" {
			fmt.Println("Error: Kategori tidak ditemukan (FK Violation)")
			return nil, domain.ErrInvalidInput
		}
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	product, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[model.Product])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		log.Println(err)
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) DeleteProduct(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	rows, err := r.DB.Query(ctx, queries.DELETE_PRODUCT, id)
	if err != nil {
		log.Println(err)
		return err
	}
	defer rows.Close()

	_, err = pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[model.Product])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.ErrNotFound
		}
		log.Println(err)
		return err
	}

	return nil
}
