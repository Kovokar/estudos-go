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

	//Camada de repository
	ProductRepository := repository.NewProductRepository(dbConnection)
	ClienteRepository := repository.NewClienteRepository(dbConnection)
	//Camada usecase
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)
	ClienteUseCase := usecase.NewClienteUseCase(ClienteRepository)
	//Camada de controllers
	ProductController := controller.NewProductController(ProductUseCase)
	ClienteController := controller.NewClienteController(ClienteUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("/product/:productId", ProductController.GetProductById)
	//--------------------------------------------------------------------------------
	server.GET("/clientes", ClienteController.GetClientes)
	server.POST("/cliente", ClienteController.CreateCliente)
	server.GET("/clientes/:clienteId", ClienteController.GetClienteById)
	//--------------------------------------------------------------------------------







	
	server.Run(":8000")

}
