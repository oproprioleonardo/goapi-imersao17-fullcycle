package webserver

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/oproprioleonardo/imersao17/goapi/internal/entity"
	"github.com/oproprioleonardo/imersao17/goapi/internal/service"
	"net/http"
)

type WebCategoryHandler struct {
	CategoryService *service.CategoryService
}

func NewCategoryHandler(categoryService *service.CategoryService) *WebCategoryHandler {
	return &WebCategoryHandler{CategoryService: categoryService}
}

func (wch *WebCategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	if categories, err := wch.CategoryService.GetCategories(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		if err := json.NewEncoder(w).Encode(categories); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (wch *WebCategoryHandler) GetCategory(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
	if category, err := wch.CategoryService.GetCategory(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		if err := json.NewEncoder(w).Encode(category); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (wch *WebCategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category entity.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if result, err := wch.CategoryService.CreateCategory(category.Name); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		if err := json.NewEncoder(w).Encode(result); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
