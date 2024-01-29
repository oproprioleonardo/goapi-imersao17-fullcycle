package webserver

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/oproprioleonardo/imersao17/goapi/internal/entity"
	"github.com/oproprioleonardo/imersao17/goapi/internal/service"
	"net/http"
)

type WebProductHandler struct {
	ProductService *service.ProductService
}

func NewWebProductHandler(productService *service.ProductService) *WebProductHandler {
	return &WebProductHandler{ProductService: productService}
}

func (wph *WebProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	if products, err := wph.ProductService.GetProducts(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		if err := json.NewEncoder(w).Encode(products); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (wph *WebProductHandler) GetProductsByCategoryID(w http.ResponseWriter, r *http.Request) {
	categoryID := chi.URLParam(r, "categoryID")
	if categoryID == "" {
		http.Error(w, "categoryID is required", http.StatusBadRequest)
		return
	}
	if products, err := wph.ProductService.GetProductsByCategoryID(categoryID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		if err := json.NewEncoder(w).Encode(products); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (wph *WebProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if result, err := wph.ProductService.CreateProduct(product.Name, product.Description, product.CategoryID, product.ImageURL, product.Price); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		if err := json.NewEncoder(w).Encode(result); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (wph *WebProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
	if product, err := wph.ProductService.GetProduct(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		if err := json.NewEncoder(w).Encode(product); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
