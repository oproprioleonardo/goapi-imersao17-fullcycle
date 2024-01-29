package service

import (
	"github.com/oproprioleonardo/imersao17/goapi/internal/database"
	"github.com/oproprioleonardo/imersao17/goapi/internal/entity"
)

type ProductService struct {
	ProductDB database.ProductDB
}

func NewProductService(db database.ProductDB) *ProductService {
	return &ProductService{ProductDB: db}
}

func (ps *ProductService) GetProducts() ([]*entity.Product, error) {
	products, err := ps.ProductDB.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *ProductService) GetProduct(id string) (*entity.Product, error) {
	product, err := ps.ProductDB.GetProduct(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (ps *ProductService) GetProductsByCategoryID(categoryID string) ([]*entity.Product, error) {
	products, err := ps.ProductDB.GetProductsByCategoryID(categoryID)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *ProductService) CreateProduct(name, description, categoryID, imageURL string, price float64) (*entity.Product, error) {
	product := entity.NewProduct(name, description, categoryID, imageURL, price)
	_, err := ps.ProductDB.CreateProduct(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}
