package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Cria um leitor a partir da entrada padrão (teclado)
	reader := bufio.NewReader(os.Stdin)

	// Loop para ler várias linhas
	for {
		// Lê uma linha de texto até encontrar uma nova linha
		input, _ := reader.ReadString('\n')

		// Remove a nova linha (\n) no final
		// input = input[:len(input)-1]

		// Se a entrada for vazia, continua o loop
		if input != "" {
			break
		}

		// Se o usuário digitar "sair", o programa encerra
		if input == "sair" {
			break
		}

		// Imprime o texto digitado
		fmt.Println(input)
	}
}
