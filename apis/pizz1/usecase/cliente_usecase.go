package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ClienteUsecase struct {
	repository repository.ClienteRepository
}

func NewClienteUseCase(repo repository.ClienteRepository) ClienteUsecase {
	return ClienteUsecase{
		repository: repo,
	}
}

func (cu *ClienteUsecase) GetClientes() ([]model.Cliente, error) {
	return cu.repository.GetClientes()
}


func (cu *ClienteUsecase) GetClienteById(id_cliente int) (*model.Cliente, error) {

	cliente, err := cu.repository.GetCleinteById(id_cliente)
	if err != nil {
		return nil, err
	}

	return cliente, nil
}


func (cu *ClienteUsecase) CreateCliente(cliente model.Cliente) (model.Cliente, error) {

	ClienteId, err := cu.repository.CreateCliente(cliente)
	if err != nil {
		return model.Cliente{}, err
	}

	cliente.ID = ClienteId

	return cliente, nil
}

