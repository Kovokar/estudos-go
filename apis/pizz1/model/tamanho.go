package model

type Tamanho struct {
	ID          int     `json:id_tamanho`
	Nome        string  `json:nome`
	Modificador float64 `json:modificador`
}
