package main

import (
	"fmt"
	"math"
	"time"
)

var voltas = 9999
var menor, consumoTotal, pessoasTotais int
var num1, num2, cont int
var slilce [][]int

func main() {
	start := time.Now()
	exec()
	elapsed := time.Since(start)
	fmt.Printf("\nTempo de execução: %s\n", elapsed)
}

func mostrar(arr [][]int) {
	cont++
	fmt.Printf("Cidade# %v:\n", cont)
	for j, sub := range arr {
		for i := 0; i < len(sub)-1; i++ {
			if j == len(arr)-1 {
				fmt.Printf("%v-%v", sub[i], sub[i+1])
			} else {
				fmt.Printf("%v-%v ", sub[i], sub[i+1])
			}
		}
	}

	consMedio := (float64(consumoTotal) / float64(pessoasTotais))
	truncated := math.Floor(consMedio*100) / 100
	fmt.Printf("\n")
	fmt.Printf("Consumo medio: %.2f m3.\n", truncated)
	fmt.Printf("\n")

}

func recebeValores(voltas int) {
	consumoTotal = 0
	pessoasTotais = 0
	for i := 0; i < voltas; i++ {
		fmt.Scanf("%v %v", &num1, &num2)

		consumoTotal += num2
		pessoasTotais += num1
		slilce = append(slilce, []int{num1, num2 / num1})
	}
}

func bubbleSort(arr [][]int) [][]int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j][1] > arr[j+1][1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func isEqual(arr [][]int) [][]int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j][1] == arr[j+1][1] {
				arr[j][0] = arr[j][0] + arr[j+1][0]
				arr = RemoveIndex(arr, j+1)
				n = len(arr)
			}
		}
	}
	return arr
}

func RemoveIndex(s [][]int, index int) [][]int {
	return append(s[:index], s[index+1:]...)
}

func esvaziaSlice(arr [][]int) [][]int {
	return arr[len(arr):]
}

func exec() {
	fmt.Scan(&voltas)
	for voltas != 0 {
		recebeValores(voltas)
		slilce = bubbleSort(slilce)
		slilce = isEqual(slilce)
		slilce = isEqual(slilce)
		slilce = isEqual(slilce)
		slilce = isEqual(slilce)
		mostrar(slilce)
		slilce = esvaziaSlice(slilce)

		// Verifica se o próximo valor de voltas é zero antes de imprimir uma linha vazia
		fmt.Scan(&voltas)
		if voltas != 0 {
			// fmt.Print("\n") // Só imprime uma linha vazia se o próximo valor de voltas não for zero
		}
	}
}
