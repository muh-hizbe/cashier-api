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

type ProductHandler struct {
	repo *repository.ProductRepository
}

func NewProductHandler(repo *repository.ProductRepository) *ProductHandler {
	return &ProductHandler{repo: repo}
}

func (h *ProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	path := strings.TrimPrefix(r.URL.Path, "/api/products")

	switch {
	case path == "" || path == "/":
		h.handleCollection(w, r)
	default:
		h.handleItem(w, r)
	}
}

func (h *ProductHandler) handleCollection(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetProducts(w, r)
	case http.MethodPost:
		h.NewProduct(w, r)
	default:
		response.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *ProductHandler) handleItem(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.GetProduct(w, r, id)
	case http.MethodPut:
		h.UpdateProduct(w, r, id)
	case http.MethodDelete:
		h.DeleteProduct(w, r, id)
	default:
		response.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.repo.GetProducts()
	if err != nil {
		response.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	response.Success(w, "Products retrieved successfully", products)
}

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request, id int) {
	product, err := h.repo.GetProduct(id)
	if err != nil {
		response.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	response.Success(w, "Product retrieved successfully", product)
}

func (h *ProductHandler) NewProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct model.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		response.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	product, err := h.repo.CreateProduct(&newProduct)
	if err != nil {
		response.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	response.Success(w, "Product stored successfully", product)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request, id int) {
	var updatedProduct model.Product
	err := json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		response.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	product, err := h.repo.UpdateProduct(&updatedProduct, id)
	if err != nil {
		response.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	response.Success(w, "Product updated successfully", product)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request, id int) {
	err := h.repo.DeleteProduct(id)
	if err != nil {
		response.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	response.Success(w, "Product deleted successfully", nil)
}
