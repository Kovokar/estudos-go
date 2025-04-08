package main

import "fmt"

func main() {
	// Declaração de um array com 5 elementos
	var arr [5]int

	// Atribuindo valores ao array
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40
	arr[4] = 50

	// Imprimindo o array
	fmt.Println("Array:", arr)

	// Acessando um valor específico
	fmt.Println("Valor na posição 2:", arr[2])
}
