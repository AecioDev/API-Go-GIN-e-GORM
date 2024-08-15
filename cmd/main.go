package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	//Camada de Repository
	ProductRepository := repository.NewProductRepository(dbConnection)

	//Camada UseCase
	ProductUsecase := usecase.NewProductUseCase(ProductRepository)

	//Camada de Controllers
	ProductController := controller.NewProductController(ProductUsecase)

	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)
	server.PUT("/product", ProductController.UpdateProduct)
	server.GET("/product/:id", ProductController.GetProductById)
	server.DELETE("/product/:id", ProductController.DeleteProductById)

	server.Run(":8000")

}
