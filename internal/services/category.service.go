package services

import (
	"github.com/muh-hizbe/cashier-api/internal/model"
	"github.com/muh-hizbe/cashier-api/internal/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetCategories() ([]model.Category, error) {
	return s.repo.GetCategories()
}

func (s *CategoryService) GetCategory(id int) (*model.Category, error) {
	return s.repo.GetCategory(id)
}

func (s *CategoryService) CreateCategory(category *model.Category) (*model.Category, error) {
	return s.repo.CreateCategory(category)
}

func (s *CategoryService) UpdateCategory(category *model.Category, id int) (*model.Category, error) {
	return s.repo.UpdateCategory(category, id)
}

func (s *CategoryService) DeleteCategory(id int) error {
	return s.repo.DeleteCategory(id)
}
