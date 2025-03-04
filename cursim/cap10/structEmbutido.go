package main

import "fmt"

type Pessoa struct {
	nome  string
	idade int64
}

type Func struct {
	pessoa  Pessoa
	funcao  string
	salario float64
}

func main() {

	pessoa1 := Pessoa{
		nome:  "felipe",
		idade: 24,
	}

	pessoa2 := Pessoa{
		nome:  "Rosario",
		idade: 27,
	}

	func1 := Func{
		pessoa:  pessoa1,
		funcao:  "99",
		salario: 3000,
	}

	fmt.Println(func1)
	fmt.Println(pessoa2)

}
