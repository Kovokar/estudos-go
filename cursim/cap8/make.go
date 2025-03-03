package main

import "fmt"

func main() {
	// Criando um slice com capacidade inicial de 5 e comprimento 0
	slice1 := make([]int, 0, 5)

	// Adicionando elementos ao slice
	slice1 = append(slice1, 1)
	slice1 = append(slice1, 2)
	slice1 = append(slice1, 3)

	fmt.Println("Slice1:", slice1)

	// Criando um slice diretamente com make, inicializado com 3 elementos
	slice2 := make([]int, 3)
	slice2[0] = 10
	slice2[1] = 20
	slice2[2] = 30

	fmt.Println("Slice2:", slice2)

	// Modificando o valor de um slice
	slice2[1] = 50
	fmt.Println("Slice2 depois da modificação:", slice2)

	// Redimensionando um slice
	slice2 = append(slice2, 40)
	fmt.Println("Slice2 após adicionar um elemento:", slice2)

	// Criando um slice de slices (para armazenar múltiplos slices)
	slice3 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("Slice3 (slice de slices):", slice3)
}
