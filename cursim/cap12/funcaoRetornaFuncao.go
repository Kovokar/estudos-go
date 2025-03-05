package main

import "fmt"

// Função que retorna duas funções
func operacoes(a int) (func() int, func() int) {
	soma := a + 10
	subtrai := a - 10

	somaFunc := func() int {
		return soma
	}

	subtraiFunc := func() int {
		return subtrai
	}

	return somaFunc, subtraiFunc
}

func main() {
	// Chamando a função que retorna duas funções
	soma, subtrai := operacoes(20)

	fmt.Println(soma())    // Output: 30
	fmt.Println(subtrai()) // Output: 10
}
