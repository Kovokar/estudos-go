package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Veiculo struct {
	ID    int32
	Name  string
	Rodas int
}

func main() {
	moto := Veiculo{1, "Moto", 2}
	carro := Veiculo{2, "Carro", 4}
	onibus := Veiculo{3, "onibus", 6}

	var veiculos = []Veiculo{moto, carro, onibus}

	b, err := json.Marshal(veiculos)

	if err != nil {
		panic(err)
	}
	fmt.Println("Com MARSHAL")
	os.Stdout.Write(b)
	fmt.Print("\n")
	// fmt.Println(string(b))

	//FORMA DOIS
	fmt.Println("Com ENCODER")
	encd := json.NewEncoder(os.Stdout)
	encd.Encode(veiculos)
}
