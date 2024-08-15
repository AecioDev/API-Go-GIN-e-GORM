package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ProductUsecase struct {
	//Repository
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) GetProductById(productId int) (*model.Product, error) {
	return pu.repository.GetProductById(productId)
}

func (pu *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {

	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId

	return product, nil
}

func (pu *ProductUsecase) UpdateProduct(product model.Product) (model.Product, error) {
	return pu.repository.UpdateProduct(product)
}

func (pu *ProductUsecase) DeleteProductById(productId int) (string, error) {
	return pu.repository.DeleteProductById(productId)
}
