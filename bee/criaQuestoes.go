package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 10; i <= 1026; i++ {
		// Cria o nome do arquivo
		fileName := fmt.Sprintf("%d.go", i)

		// Cria o arquivo
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Println("Erro ao criar o arquivo:", err)
			continue
		}
		defer file.Close()

		// Escreve algo no arquivo (opcional)
		file.WriteString("// Arquivo gerado automaticamente: " + fileName + "\n")

		// Confirmação de criação
		fmt.Println("Arquivo criado:", fileName)
	}
}
