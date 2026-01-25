package repository

import (
	"errors"

	"github.com/muh-hizbe/cashier-api/internal/model"
)

type ProductRepository struct {
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

func (r *ProductRepository) GetProducts() ([]model.Product, error) {
	return model.GetProducts(), nil
}

func (r *ProductRepository) GetProduct(id int) (*model.Product, error) {
	for _, p := range model.Products {
		if p.ID == id {
			return &p, nil
		}
	}

	return nil, errors.New("product not found")
}

func (r *ProductRepository) CreateProduct(product *model.Product) (model.Product, error) {
	product.ID = len(model.Products) + 1
	model.Products = append(model.Products, *product)
	return *product, nil
}

func (r *ProductRepository) UpdateProduct(product *model.Product, id int) (model.Product, error) {
	for i, p := range model.Products {
		if p.ID == id {
			product.ID = id
			model.Products[i] = *product
			return *product, nil
		}
	}

	return *product, errors.New("product not found")
}

func (r *ProductRepository) DeleteProduct(id int) error {
	for i, p := range model.Products {
		if p.ID == id {
			model.Products = append(model.Products[:i], model.Products[i+1:]...)
			return nil
		}
	}

	return errors.New("product not found")
}
