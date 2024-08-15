package repository

import (
	"fmt"
	"go-api/model"

	"gorm.io/gorm"
)

type ProductRepository struct {
	connection *gorm.DB
}

func NewProductRepository(connection *gorm.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {

	var productList []model.Product

	result := pr.connection.Find(&productList)

	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return []model.Product{}, result.Error
	}

	return productList, nil
}

func (pr *ProductRepository) GetProductById(produtoId int) (*model.Product, error) {

	var productObj model.Product

	result := pr.connection.First(&productObj, produtoId)

	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return &model.Product{}, result.Error
	}

	return &productObj, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {

	result := pr.connection.Create(&product)

	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return 0, result.Error
	}

	return product.ID, nil
}

func (pr *ProductRepository) UpdateProduct(product model.Product) (model.Product, error) {

	result := pr.connection.Save(&product)

	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return product, result.Error
	}

	return product, nil
}

func (pr *ProductRepository) DeleteProductById(productId int) (string, error) {

	result := pr.connection.Delete(model.Product{}, productId)

	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return "Erro ao Deletar o Produto:\n", result.Error
	}

	return "Produto Deletado com Sucesso!", nil
}
