package main

import "fmt"

// Definição de um tipo struct
type Pessoa struct {
	Nome  string
	Idade int
}

// Método associado à struct Pessoa
func (p Pessoa) Saudacao() {
	fmt.Printf("Olá, meu nome é %s e tenho %d anos.\n", p.Nome, p.Idade)
}

func (p *Pessoa) Envelhecer() {
	p.Idade++
}

func main() {
	pessoa := Pessoa{"Carlos", 30}
	pessoa.Saudacao() // Chamando o método
	pessoa.Envelhecer()
	pessoa.Saudacao()         // Chamando o método
	fmt.Println(pessoa.Idade) // 31

}
