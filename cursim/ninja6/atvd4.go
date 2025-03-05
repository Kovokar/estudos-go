package main

import "fmt"

type Pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p Pessoa) infos() {
	fmt.Println("meu nome é:", p.nome, p.sobrenome, "e tenho", p.idade, "anos")
}

func main() {
	pedro := Pessoa{"Pedro", "Guilherme", 18}
	pedro.infos()

}
