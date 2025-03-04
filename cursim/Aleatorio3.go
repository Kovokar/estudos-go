package main

import (
	"fmt"
	"math"
)

// Função para verificar se um número é primo
func isPrimo(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Função que calcula a soma dos números primos em uma lista
func somaPrimos(nums []int) int {
	soma := 0
	for _, num := range nums {
		if isPrimo(num) {
			soma += num
		}
	}
	return soma
}

func main() {
	// Testando a função
	entrada := []int{2, 3, 4, 5, 10, 11, 13, 20}
	resultado := somaPrimos(entrada)
	fmt.Printf("A soma dos números primos na lista é: %d\n", resultado)
}
