package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, caso int

	for {
		_, err := fmt.Scan(&n)
		if err != nil {
			break
		}

		caso++
		totalNumeros := 1 + (n * (n + 1) / 2) // Fórmula para o tamanho total da sequência
		var builder strings.Builder
		builder.WriteString("0")

		for i := 1; i <= n; i++ {
			for j := 0; j < i; j++ {
				builder.WriteString(fmt.Sprintf(" %d", i))
			}
		}

		// Tratamento da palavra "numero" ou "numeros"
		if totalNumeros == 1 {
			fmt.Printf("Caso %d: %d numero\n", caso, totalNumeros)
		} else {
			fmt.Printf("Caso %d: %d numeros\n", caso, totalNumeros)
		}

		// Imprime a sequência completa de uma vez só
		fmt.Println(builder.String())

		// Linha em branco após cada caso
		fmt.Println()
	}
}
