package main

import "fmt"

type Veiculo struct {
	cor    string
	portas int
}

type Caminhonete struct {
	veiculo     Veiculo
	quatroRodas bool
}

type Sedan struct {
	veiculo    Veiculo
	modeloLuxo bool
}

func main() {

	vPreto := Veiculo{
		cor:    "preto",
		portas: 4,
	}

	golAzulBolinha := Veiculo{
		cor:    "Azul",
		portas: 2,
	}

	caminhonete1 := Caminhonete{
		veiculo:     vPreto,
		quatroRodas: vPreto.portas == 4,
	}

	caminhonete2 := Caminhonete{
		veiculo:     golAzulBolinha,
		quatroRodas: golAzulBolinha.portas == 4,
	}

	seda1 := Sedan{golAzulBolinha, false}

	fmt.Println(caminhonete1)
	fmt.Println(caminhonete2)
	fmt.Println(seda1)

}
