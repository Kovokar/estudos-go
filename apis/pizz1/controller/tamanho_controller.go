package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type tamanhoController struct {
	tamanhoUseCase usecase.TamanhoUsecase
}

func NewTamanhoController(usecase usecase.TamanhoUsecase) tamanhoController {
	return tamanhoController{
		tamanhoUseCase: usecase,
	}
}

func (c *tamanhoController) GetTamanhos(ctx *gin.Context) {
	tamanhos, err := c.tamanhoUseCase.GetTamanhos()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, tamanhos)
}

func (c *tamanhoController) CreateTamanho(ctx *gin.Context) {
	var tamanho model.Tamanho
	err := ctx.BindJSON(&tamanho)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedTamanho, err := c.tamanhoUseCase.CreateTamanho(tamanho)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedTamanho)
}

func (c *tamanhoController) GetTamanhoById(ctx *gin.Context) {
	id := ctx.Param("tamanhoId")
	if id == "" {
		response := model.Response{
			Message: "Id do tamanho não pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	tamanhoId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Id do tamanho precisa ser um número",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	tamanho, err := c.tamanhoUseCase.GetTamanhoById(tamanhoId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if tamanho == nil {
		response := model.Response{
			Message: "Tamanho não encontrado",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, tamanho)
}
