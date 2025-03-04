package main

import (
	"fmt"
	"sort"
	"strings"
)

// Função que recebe uma string com números e retorna o maior número formado pelos dígitos
func maiorNumeroFormado(s string) string {
	// Cria uma slice para armazenar todos os dígitos
	var digitos []rune
	
	// Itera sobre cada número na string
	numeros := strings.Fields(s)
	for _, numero := range numeros {
		// Adiciona os dígitos de cada número à slice
		for _, digit := range numero {
			digitos = append(digitos, digit)
		}
	}
	
	// Ordena os dígitos em ordem decrescente
	sort.Sort(sort.Reverse(sort.StringSlice(digitos)))
	
	// Converte a slice de runes de volta para uma string
	return string(digitos)
}

func main() {
	// Testando a função
	entrada := "123 456 789"
	resultado := maiorNumeroFormado(entrada)
	fmt.Printf("O maior número possível formado é: %s\n", resultado)
}
