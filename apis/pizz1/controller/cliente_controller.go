package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type clienteController struct {
	clienteUseCase usecase.ClienteUsecase
}

func NewClienteController(usecase usecase.ClienteUsecase) clienteController {
	return clienteController{
		clienteUseCase: usecase,
	}
}

func (c *clienteController) GetClientes(ctx *gin.Context) {

	clientes, err := c.clienteUseCase.GetClientes()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, clientes)
}

func (c *clienteController) CreateCliente(ctx *gin.Context) {

	var cliente model.Cliente
	err := ctx.BindJSON(&cliente)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedCliente, err := c.clienteUseCase.CreateCliente(cliente)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedCliente)
}

func (c *clienteController) GetClienteById(ctx *gin.Context) {

	id := ctx.Param("clienteId")
	if id == "" {
		response := model.Response{
			Message: "Id do cliente nao pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	clienteId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Id do cliente precisa ser um numero",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	cliente, err := c.clienteUseCase.GetClienteById(clienteId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if cliente == nil {
		response := model.Response{
			Message: "Cliente nao foi encontrado na base de dados",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, cliente)
}
