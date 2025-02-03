package model

type Pizza struct {
	ID        int    `json:id_pizza`
	Sabor     string `json:sabor_pizza`
	PrecoBase float64 `json:preco_base_pizza`
}
