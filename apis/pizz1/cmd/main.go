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
	PizzaRepository := repository.NewPizzaRepository(dbConnection)
	TamanhoRepository := repository.NewTamanhoRepository(dbConnection)
	
	//Camada usecase
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)
	ClienteUseCase := usecase.NewClienteUseCase(ClienteRepository)
	PizzaUsecase := usecase.NewPizzaUsecase(PizzaRepository)
	TamanhoUsecase := usecase.NewTamanhoUsecase(TamanhoRepository)
	
	//Camada de controllers
	ProductController := controller.NewProductController(ProductUseCase)
	ClienteController := controller.NewClienteController(ClienteUseCase)
	pizzaController := controller.NewPizzaController(PizzaUsecase)
	tamanhoController := controller.NewTamanhoController(TamanhoUsecase)
	

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
	server.GET("/pizzas", pizzaController.GetPizzas)
	server.POST("/pizza", pizzaController.CreatePizza)
	server.GET("/pizzas/:clienteId", pizzaController.GetPizzaById)
	//--------------------------------------------------------------------------------
	server.GET("/tamanhos", tamanhoController.GetTamanhos)
	server.POST("/tamanho", tamanhoController.CreateTamanho)
	server.GET("/tamanhos/:tamanhoId", tamanhoController.GetTamanhoById)
	//--------------------------------------------------------------------------------
	
	server.Run(":8000")

}
