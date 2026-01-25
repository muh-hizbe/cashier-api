package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/muh-hizbe/cashier-api/internal/model"
	"github.com/muh-hizbe/cashier-api/internal/repository"
	"github.com/muh-hizbe/cashier-api/internal/response"
)

type CategoryHandler struct {
	repo *repository.CategoryRepository
}

func NewCategoryHandler(repo *repository.CategoryRepository) *CategoryHandler {
	return &CategoryHandler{repo: repo}
}

func (h *CategoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	path := strings.TrimPrefix(r.URL.Path, "/api/categories")

	switch {
	case path == "" || path == "/":
		h.handleCollection(w, r)
	default:
		h.handleItem(w, r)
	}
}

func (h *CategoryHandler) handleCollection(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetCategories(w, r)
	case http.MethodPost:
		h.NewCategory(w, r)
	default:
		response.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *CategoryHandler) handleItem(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.GetCategory(w, r, id)
	case http.MethodPut:
		h.UpdateCategory(w, r, id)
	case http.MethodDelete:
		h.DeleteCategory(w, r, id)
	default:
		response.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.repo.GetCategories()
	if err != nil {
		response.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	response.Success(w, "Categories retrieved successfully", categories)
}

func (h *CategoryHandler) GetCategory(w http.ResponseWriter, r *http.Request, id int) {
	category, err := h.repo.GetCategory(id)
	if err != nil {
		response.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	response.Success(w, "Category retrieved successfully", category)
}

func (h *CategoryHandler) NewCategory(w http.ResponseWriter, r *http.Request) {
	var newCategory model.Category
	err := json.NewDecoder(r.Body).Decode(&newCategory)
	if err != nil {
		response.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	category, err := h.repo.CreateCategory(&newCategory)
	if err != nil {
		response.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	response.Success(w, "Category stored successfully", category)
}

func (h *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request, id int) {
	var updatedCategory model.Category
	err := json.NewDecoder(r.Body).Decode(&updatedCategory)
	if err != nil {
		response.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	category, err := h.repo.UpdateCategory(&updatedCategory, id)
	if err != nil {
		response.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	response.Success(w, "Category updated successfully", category)
}

func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request, id int) {
	err := h.repo.DeleteCategory(id)
	if err != nil {
		response.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	response.Success(w, "Category deleted successfully", nil)
}
