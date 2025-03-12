package main

import (
	"fmt"
)

func calcularTotal(numeros []int) int {
	total := 0
	for _, num := range numeros {
		total += num
	}
	return total
}

func get() {
	var stars int
	var aux bool = false
	fmt.Scan(&stars)

	numeros := make([]int, stars)

	for i := 0; i < stars; i++ {
		fmt.Scan(&numeros[i])
	}

	tot := calcularTotal(numeros)
	var count int // Variável para contar as iterações do loop i

	for i := 0; i < len(numeros); i++ {
		count++ // Incrementa a cada iteração do loop
		if numeros[i] == 0 {
			break
		}
		if numeros[i]%2 == 0 {
			for j := i; j > 0; j-- {
				numeros[j] -= 1
				if numeros[j] == 0 || numeros[j-1] == 0 {
					aux = true
					break
				}
			}
		}
		if aux {
			break
		}
		numeros[i] -= 1
		fmt.Println(numeros)
	}

	newTot := calcularTotal(numeros)
	fmt.Printf("%v %v", count, tot-(tot-newTot))
}

func main() {
	get()
}
