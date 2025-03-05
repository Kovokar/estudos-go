package main

import "fmt"

type Pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func mudeMe(p *Pessoa) {
	fmt.Println(p.idade)
	p.idade++
	fmt.Println(p.idade)

}

func changeMe(p Pessoa) {
	fmt.Println(p.idade)
	p.idade++
	fmt.Println(p.idade)
}

func main() {
	pedro := Pessoa{"Pedro", "Guilherme", 18}
	mudeMe(&pedro)
	// changeMe(pedro)

}
