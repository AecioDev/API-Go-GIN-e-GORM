package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	//Usecase
	productUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		productUsecase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {

	products, err := p.productUsecase.GetProducts()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *productController) GetProductById(ctx *gin.Context) {

	id := ctx.Param("id")
	if id == "" {
		response := model.Response{
			Message: "Id do Produto não pode ser nulo!",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Id do Produto precisa ser um número!",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.productUsecase.GetProductById(productId)

	if err != nil {
		response := model.Response{
			Message: "Erro ao obter o produto: " + err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	if product == nil {
		response := model.Response{
			Message: "Produto não foi encontrado na base de dados!",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func (p *productController) CreateProduct(ctx *gin.Context) {

	var product model.Product
	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.productUsecase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (p *productController) UpdateProduct(ctx *gin.Context) {

	var product model.Product
	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	updatedProduct, err := p.productUsecase.UpdateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, updatedProduct)
}

func (p *productController) DeleteProductById(ctx *gin.Context) {

	id := ctx.Param("id")
	if id == "" {
		response := model.Response{
			Message: "Id do Produto não pode ser nulo!",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Id do Produto precisa ser um número!",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.productUsecase.GetProductById(productId)

	if err != nil {
		response := model.Response{
			Message: "Erro ao obter o produto: " + err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	if product == nil {
		response := model.Response{
			Message: "Produto não foi encontrado na base de dados!",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	} else {
		result, err := p.productUsecase.DeleteProductById(productId)
		if err != nil {
			response := model.Response{
				Message: "Erro interno ao Deletar o Produto:\n" + err.Error(),
			}
			ctx.JSON(http.StatusInternalServerError, response)
			return
		}

		response := model.Response{
			Message: result,
		}
		ctx.JSON(http.StatusOK, response)
	}
}
