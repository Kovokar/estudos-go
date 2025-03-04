package main

import (
	"errors"
	"fmt"
	"sort"
)

// Função que retorna o segundo maior número único de uma lista de inteiros
func segundoMaior(nums []int) (int, error) {
	// Cria um mapa para garantir que os números são únicos
	numMap := make(map[int]struct{})

	// Adiciona os números ao mapa para garantir a unicidade
	for _, num := range nums {
		numMap[num] = struct{}{}
	}

	// Se houver menos de 2 números distintos, retorna erro
	if len(numMap) < 2 {
		return 0, errors.New("não há números suficientes para calcular o segundo maior")
	}

	// Cria uma slice com os números únicos
	var uniqueNums []int
	for num := range numMap {
		uniqueNums = append(uniqueNums, num)
	}

	// Ordena os números de forma crescente
	sort.Ints(uniqueNums)

	// Retorna o segundo maior número
	return uniqueNums[len(uniqueNums)-2], nil
}

func main() {
	// Testando a função
	entrada := []int{3, 1, 4, 5, 2}
	resultado, err := segundoMaior(entrada)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("O segundo maior número é: %d\n", resultado)
	}
}
