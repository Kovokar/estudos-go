// main.go
package main

import (
	"fmt"
	"log"
	"pizza-api/controllers"
	"pizza-api/dbs" // Aqui você importa o pacote de configuração do banco
	"pizza-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializa a conexão com o banco de dados
	config := dbs.DBConfig{
		User:     "postgres",
		Password: "Girassol0@",
		DbName:   "gepeto",
		Host:     "localhost", // Host definido como localhost
		Port:     "5432",      // Porta padrão do PostgreSQL
		SSLMode:  "disable",   // SSL desabilitado
	}

	dbs.InitDB(config)

	// Usando AutoMigrate para criar as tabelas automaticamente
	if err := dbs.GetDB().AutoMigrate(&models.Cliente{}, &models.Pizza{}, &models.Tamanho{}, &models.Pedido{}, &models.ItensPedido{}); err != nil {
		log.Fatal("Erro ao migrar os modelos:", err)
	}

	fmt.Println("Tabelas criadas com sucesso!")

	// Aqui você pode continuar configurando o Gin e as rotas
	router := gin.Default()

	// Roteamento
	router.POST("/clientes", controllers.CreateCliente)
	router.GET("/clientes", controllers.GetClientes)
	router.POST("/pizzas", controllers.CreatePizza)
	router.GET("/pizzas", controllers.GetPizzas)
	router.POST("/tamanhos", controllers.CreateTamanho)
	router.GET("/tamanhos", controllers.GetTamanhos)
	router.POST("/pedidos", controllers.CreatePedido)
	router.GET("/pedidos", controllers.GetPedidos)
	router.POST("/itenspedidos", controllers.CreateItemPedido)
	router.GET("/itenspedidos", controllers.GetItensPedidos)

	// Inicia o servidor na porta 8080
	router.Run(":8080")
}
