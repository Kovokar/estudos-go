package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type ClienteRepository struct {
	connection *sql.DB
}

func NewClienteRepository(connection *sql.DB) ClienteRepository {
	return ClienteRepository{
		connection: connection,
	}
}


func (cr *ClienteRepository) GetClientes() ([]model.Cliente, error) {

	query := "SELECT id, nome, endereco, telefone, bairro FROM cliente"
	rows, err := cr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Cliente{}, err
	}

	var clienteList []model.Cliente
	var clienteObj model.Cliente

	for rows.Next() {
		err = rows.Scan(
			&clienteObj.ID,
			&clienteObj.Nome,
			&clienteObj.Endereco,
			&clienteObj.Tel,
			&clienteObj.Bairro,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Cliente{}, err
		}

		clienteList = append(clienteList, clienteObj)
	}

	rows.Close()

	return clienteList, nil
}


func (cu *ClienteRepository) GetCleinteById(id_cliente int) (*model.Cliente, error) {

	query, err := cu.connection.Prepare("SELECT * FROM cliente WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var cliente model.Cliente

	err = query.QueryRow(id_cliente).Scan(
		&cliente.ID,
		&cliente.Nome,
		&cliente.Endereco,
		&cliente.Tel,
		&cliente.Bairro,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()
	return &cliente, nil
}


func (cr *ClienteRepository) CreateCliente(cliente model.Cliente) (int, error) {

	var id int
	query, err := cr.connection.Prepare("INSERT INTO cliente" +
		"(nome, endereco, telefone, bairro)" +
		" VALUES ($1, $2, $3, $4) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(cliente.Nome, cliente.Endereco, cliente.Tel, cliente.Bairro).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}

