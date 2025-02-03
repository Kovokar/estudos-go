package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type pizzaController struct {
	pizzaUseCase usecase.PizzaUsecase
}

func NewPizzaController(usecase usecase.PizzaUsecase) pizzaController {
	return pizzaController{
		pizzaUseCase: usecase,
	}
}

func (c *pizzaController) GetPizzas(ctx *gin.Context) {

	pizzas, err := c.pizzaUseCase.GetPizzas()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, pizzas)
}

func (c *pizzaController) CreatePizza(ctx *gin.Context) {

	var pizza model.Pizza
	err := ctx.BindJSON(&pizza)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedPizza, err := c.pizzaUseCase.CreatePizza(pizza)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedPizza)
}

func (c *pizzaController) GetPizzaById(ctx *gin.Context) {

	id := ctx.Param("pizzaId")
	if id == "" {
		response := model.Response{
			Message: "Id da pizza não pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	pizzaId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Id da pizza precisa ser um número",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	pizza, err := c.pizzaUseCase.GetPizzaById(pizzaId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if pizza == nil {
		response := model.Response{
			Message: "Pizza não foi encontrada na base de dados",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, pizza)
}
