package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ProductUseCase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		repository: repo,
	}
}

func (p *ProductUseCase) GetProducts() ([]model.Product, error) {
	return p.repository.GetProducts()
}

func (p *ProductUseCase) CreateProduct(product model.Product) (model.Product, error) {

	productId, err := p.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}
	product.ID = productId
	return product, nil
}

func (p *ProductUseCase) GetProductById(id int) (*model.Product, error) {
	return p.repository.GetProductById(id)
}