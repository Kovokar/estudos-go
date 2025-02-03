package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type PizzaRepository struct {
	connection *sql.DB
}

func NewPizzaRepository(connection *sql.DB) PizzaRepository {
	return PizzaRepository{
		connection: connection,
	}
}

func (pzr *PizzaRepository) GetPizzas() ([]model.Pizza, error) {

	query := "SELECT id, sabor, precobase FROM pizza"
	rows, err := pzr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Pizza{}, err
	}

	var pizzaList []model.Pizza
	var pizzaObj model.Pizza

	for rows.Next() {
		err = rows.Scan(
			&pizzaObj.ID,
			&pizzaObj.Sabor,
			&pizzaObj.PrecoBase)

		if err != nil {
			fmt.Println(err)
			return []model.Pizza{}, err
		}

		pizzaList = append(pizzaList, pizzaObj)
	}

	rows.Close()

	return pizzaList, nil
}

func (pzr *PizzaRepository) CreatePizza(pizza model.Pizza) (int, error) {

	var id int
	query, err := pzr.connection.Prepare("INSERT INTO pizza" +
		"(sabor, precobase)" +
		" VALUES ($1, $2) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(pizza.Sabor, pizza.PrecoBase).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}

func (pzr *PizzaRepository) GetPizzaById(id_pizza int) (*model.Pizza, error) {

	query, err := pzr.connection.Prepare("SELECT * FROM pizza WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var pizza model.Pizza

	err = query.QueryRow(id_pizza).Scan(
		&pizza.ID,
		&pizza.Sabor,
		&pizza.PrecoBase,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()
	return &pizza, nil
}
