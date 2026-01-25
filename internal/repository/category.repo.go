package repository

import (
	"errors"

	"github.com/muh-hizbe/cashier-api/internal/model"
)

type CategoryRepository struct {
}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{}
}

func (r *CategoryRepository) GetCategories() ([]model.Category, error) {
	return model.GetCategories(), nil
}

func (r *CategoryRepository) GetCategory(id int) (*model.Category, error) {
	for _, c := range model.Categories {
		if c.ID == id {
			return &c, nil
		}
	}

	return nil, errors.New("category not found")
}

func (r *CategoryRepository) CreateCategory(category *model.Category) (model.Category, error) {
	category.ID = len(model.Categories) + 1
	model.Categories = append(model.Categories, *category)
	return *category, nil
}

func (r *CategoryRepository) UpdateCategory(category *model.Category, id int) (model.Category, error) {
	for i, c := range model.Categories {
		if c.ID == id {
			category.ID = id
			model.Categories[i] = *category
			return *category, nil
		}
	}

	return *category, errors.New("category not found")
}

func (r *CategoryRepository) DeleteCategory(id int) error {
	for i, c := range model.Categories {
		if c.ID == id {
			model.Categories = append(model.Categories[:i], model.Categories[i+1:]...)
			return nil
		}
	}

	return errors.New("category not found")
}
