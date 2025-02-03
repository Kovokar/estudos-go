package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type PizzaUsecase struct {
	repository repository.PizzaRepository
}

func NewPizzaUsecase(repo repository.PizzaRepository) PizzaUsecase {
	return PizzaUsecase{
		repository: repo,
	}
}

func (pu *PizzaUsecase) GetPizzas() ([]model.Pizza, error) {
	return pu.repository.GetPizzas()
}

func (pu *PizzaUsecase) CreatePizza(pizza model.Pizza) (model.Pizza, error) {
	// Chama o repositório para criar a pizza no banco de dados
	pizzaId, err := pu.repository.CreatePizza(pizza)
	if err != nil {
		return model.Pizza{}, err
	}

	// Atualiza o ID da pizza criada
	pizza.ID = pizzaId

	return pizza, nil
}

func (pu *PizzaUsecase) GetPizzaById(id_pizza int) (*model.Pizza, error) {
	// Chama o repositório para pegar a pizza pelo ID
	pizza, err := pu.repository.GetPizzaById(id_pizza)
	if err != nil {
		return nil, err
	}

	return pizza, nil
}
