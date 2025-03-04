package main

import "fmt"

// Função variádica que soma todos os números passados como argumento
func somar(numeros ...int) int {
	soma := 0
	for _, num := range numeros {
		soma += num
	}
	return soma
}
func exemplo(a int, b int, c int) {
	fmt.Println(a, b, c)
}

// Passa os elementos de 's' como parâmetros para a função.

func main() {
	// Chamando a função variádica
	resultado := somar(1, 2, 3, 4, 5)
	fmt.Println("Resultado da soma:", resultado) // Saída: Resultado da soma: 15

	// Também podemos passar um slice como argumento
	numeros := []int{10, 20, 30}
	resultado2 := somar(numeros...)                         // Descompactando o slice para passar os elementos individualmente
	fmt.Println("Resultado da soma com slice:", resultado2) // Saída: Resultado da soma com slice: 60

	// s := []int{1, 2, 3}
	// exemplo(s...)

}
