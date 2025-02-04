package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type TamanhoRepository struct {
	connection *sql.DB
}

func NewTamanhoRepository(connection *sql.DB) TamanhoRepository {
	return TamanhoRepository{
		connection: connection,
	}
}

func (tr *TamanhoRepository) GetTamanhos() ([]model.Tamanho, error) {
	query := "SELECT id, nome, modificador FROM tamanho"
	rows, err := tr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Tamanho{}, err
	}

	var tamanhoList []model.Tamanho
	var tamanhoObj model.Tamanho

	for rows.Next() {
		err = rows.Scan(
			&tamanhoObj.ID,
			&tamanhoObj.Nome,
			&tamanhoObj.Modificador)

		if err != nil {
			fmt.Println(err)
			return []model.Tamanho{}, err
		}

		tamanhoList = append(tamanhoList, tamanhoObj)
	}

	rows.Close()
	return tamanhoList, nil
}

func (tr *TamanhoRepository) CreateTamanho(tamanho model.Tamanho) (int, error) {
	var id int
	query, err := tr.connection.Prepare("INSERT INTO tamanho (nome, modificador) VALUES ($1, $2) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(tamanho.Nome, tamanho.Modificador).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}

func (tr *TamanhoRepository) GetTamanhoById(id_tamanho int) (*model.Tamanho, error) {
	query, err := tr.connection.Prepare("SELECT * FROM tamanho WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var tamanho model.Tamanho
	err = query.QueryRow(id_tamanho).Scan(
		&tamanho.ID,
		&tamanho.Nome,
		&tamanho.Modificador,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	query.Close()
	return &tamanho, nil
}
