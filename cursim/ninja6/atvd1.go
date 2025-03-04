package main

import "fmt"

type Pessoa struct {
	nome             string
	sobrenome        string
	saboresDeSorvete []string
}

func main() {
	pessoa1 := Pessoa{
		nome:             "Pedro",
		sobrenome:        "guilherme",
		saboresDeSorvete: []string{"kiwi", "chocomenta"},
	}
	pessoa2 := Pessoa{
		nome:             "Alessia",
		sobrenome:        "Vitória",
		saboresDeSorvete: []string{"chocolate", "Morango"},
	}

	fmt.Println(pessoa1)
	fmt.Println(pessoa2)
}
