package main

import (
	"fmt"
	"math"
	"sort"
	"time"
)

func main() {
	start := time.Now()

	var cidadeNumero int
	var N int
	var totalPessoas, consumoTotal int
	var consumoPorPessoa map[int]int

	cidadeNumero = 1

	fmt.Scan(&N)

	for N != 0 {
		totalPessoas = 0
		consumoTotal = 0
		consumoPorPessoa = make(map[int]int)

		for i := 0; i < N; i++ {
			var X, Y int
			fmt.Scan(&X, &Y)

			consumoPorPessoaPorPessoa := Y / X
			consumoPorPessoa[consumoPorPessoaPorPessoa] += X
			totalPessoas += X
			consumoTotal += Y
		}

		fmt.Printf("Cidade# %d:\n", cidadeNumero)
		keys := []int{}
		for k := range consumoPorPessoa {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		for i, consumo := range keys {
			if i == len(keys)-1 {
				fmt.Printf("%d-%d", consumoPorPessoa[consumo], consumo)
			} else {
				fmt.Printf("%d-%d ", consumoPorPessoa[consumo], consumo)
			}
		}

		consumoMedio := float64(consumoTotal) / float64(totalPessoas)
		consumoMedioArredondado := math.Floor(consumoMedio*100) / 100

		fmt.Printf("\nConsumo medio: %.2f m3.", consumoMedioArredondado)
		fmt.Scan(&N)
		if N != 0 {
			fmt.Println()
		}
		cidadeNumero++
		if N != 0 {
			fmt.Println()
		}
	}

	elapsed := time.Since(start)

	// Exibe o tempo de execução
	fmt.Printf("\nTempo de execução: %s\n", elapsed)
}
