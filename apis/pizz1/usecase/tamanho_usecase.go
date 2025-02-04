package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type TamanhoUsecase struct {
	repository repository.TamanhoRepository
}

func NewTamanhoUsecase(repo repository.TamanhoRepository) TamanhoUsecase {
	return TamanhoUsecase{
		repository: repo,
	}
}

func (tu *TamanhoUsecase) GetTamanhos() ([]model.Tamanho, error) {
	return tu.repository.GetTamanhos()
}

func (tu *TamanhoUsecase) CreateTamanho(tamanho model.Tamanho) (model.Tamanho, error) {
	// Chama o repositório para criar o tamanho no banco de dados
	tamanhoId, err := tu.repository.CreateTamanho(tamanho)
	if err != nil {
		return model.Tamanho{}, err
	}

	// Atualiza o ID do tamanho criado
	tamanho.ID = tamanhoId

	return tamanho, nil
}

func (tu *TamanhoUsecase) GetTamanhoById(id_tamanho int) (*model.Tamanho, error) {
	// Chama o repositório para pegar o tamanho pelo ID
	tamanho, err := tu.repository.GetTamanhoById(id_tamanho)
	if err != nil {
		return nil, err
	}

	return tamanho, nil
}
