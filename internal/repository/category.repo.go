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

type CategoryRepository struct {
	DB *pgxpool.Pool
}

func NewCategoryRepository(db *pgxpool.Pool) *CategoryRepository {
	return &CategoryRepository{
		DB: db,
	}
}

func (r *CategoryRepository) GetCategories() ([]model.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	rows, err := r.DB.Query(ctx, queries.GET_CATEGORIES)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	categories, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[model.Category])
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return categories, nil
}

func (r *CategoryRepository) GetCategory(id int) (*model.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	rows, err := r.DB.Query(ctx, queries.GET_CATEGORY_BY_ID, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	category, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[model.Category])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		log.Println(err)
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) CreateCategory(payload *model.Category) (*model.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	rows, err := r.DB.Query(ctx, queries.INSERT_CATEGORY, payload.Name, payload.Description)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	category, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[model.Category])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		log.Println(err)
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) UpdateCategory(payload *model.Category, id int) (*model.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	rows, err := r.DB.Query(ctx, queries.UPDATE_CATEGORY, id, payload.Name, payload.Description)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	category, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[model.Category])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		log.Println(err)
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) DeleteCategory(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	rows, err := r.DB.Query(ctx, queries.DELETE_CATEGORY, id)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23503" {
			fmt.Println("Error: Kategori tidak bisa dihapus, karena produk masih terhubung (FK Violation)")
			return domain.ErrInternal
		}
		log.Println(err)
		return err
	}
	defer rows.Close()

	_, err = pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[model.Category])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.ErrNotFound
		}
		log.Println(err)
		return err
	}

	return nil
}
