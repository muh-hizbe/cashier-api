package services

import (
	"github.com/muh-hizbe/cashier-api/internal/model"
	"github.com/muh-hizbe/cashier-api/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetProducts() ([]model.Product, error) {
	return s.repo.GetProducts()
}

func (s *ProductService) GetProduct(id int) (*model.Product, error) {
	return s.repo.GetProduct(id)
}

func (s *ProductService) CreateProduct(product *model.Product) (*model.Product, error) {
	return s.repo.CreateProduct(product)
}

func (s *ProductService) UpdateProduct(product *model.Product, id int) (*model.Product, error) {
	return s.repo.UpdateProduct(product, id)
}

func (s *ProductService) DeleteProduct(id int) error {
	return s.repo.DeleteProduct(id)
}
