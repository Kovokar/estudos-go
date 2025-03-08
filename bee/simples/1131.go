package main

import (
	"fmt"
)

func get() (float64, float64) {
	var v1, v2 float64
	fmt.Scan(&v1, &v2)
	return v1, v2
}

func main() {
	var vitoria1, vitoria2, empate int
	for {
		v1, v2 := get()
		if v1 > v2 {
			vitoria1++
		} else if v2 > v1 {
			vitoria2++
		} else {
			empate++
		}

		fmt.Println("Novo grenal (1-sim 2-nao)")
		var i int
		fmt.Scan(&i)
		if i == 2 {
			break
		}
	}

	fmt.Printf("%d grenais\n", vitoria1+vitoria2+empate)
	fmt.Printf("Inter:%d\n", vitoria1)
	fmt.Printf("Grêmio:%d\n", vitoria2)
	fmt.Printf("Empates:%d\n", empate)

	if vitoria1 > vitoria2 {
		fmt.Println("Inter venceu mais")
	} else if vitoria2 > vitoria1 {
		fmt.Println("Gremio venceu mais")
	} else {
		fmt.Println("Nao houve vencedor")
	}
}
