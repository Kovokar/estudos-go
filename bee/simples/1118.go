package main

import "fmt"

func nota_valida(nota float64) bool {
	return 0 <= nota && nota <= 10
}

func get_nota() float64 {
	for {
		var nota float64
		fmt.Scan(&nota)
		if nota_valida(nota) {
			return nota
		}
		fmt.Println("nota invalida")
	}
}

func new_calc() bool {
	for {
		fmt.Println("novo calculo (1-sim 2-nao)")
		var op float64
		fmt.Scan(&op)
		if op == 1 {
			return true
		} else if op == 2 {
			return false
		}
	}
}

func main() {
	for {
		n1 := get_nota()
		n2 := get_nota()
		med := (n1 + n2) / 2
		fmt.Printf("media = %.2f\n", med)
		if !new_calc() {
			break
		}
	}
}
