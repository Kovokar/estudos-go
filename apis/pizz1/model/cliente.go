package model

type Cliente struct {
	ID       int    `json:id_cliente`
	Nome     string `json:nome_cliente`
	Endereco string `json:endereco_cliente`
	Tel      string `json:tel_clinte`
	Bairro   string `json:bairro_cliente`
}
