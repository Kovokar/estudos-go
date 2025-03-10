package main

import (
	"fmt"
)

func getBin() {
	aux := 0
	for {
		var caw string
		fmt.Scanln(&caw)

		if caw[0] == 'c' {
			fmt.Println(aux)
			aux = 0
			break
		}

		num := ""
		for _, i := range caw {
			if i == '-' || i == ' ' {
				num += "0"
			}
			if i == '*' {
				num += "1"
			}
		}

		// Converte a string binária em um número inteiro
		value, err := binaryToInt(num)
		if err != nil {
			fmt.Println("Erro ao converter binário:", err)
			return
		}

		aux += value
	}
}

func binaryToInt(bin string) (int, error) {
	// Converte a string binária para inteiro
	var result int
	for _, char := range bin {
		result *= 2
		if char == '1' {
			result += 1
		}
	}
	return result, nil
}

func main() {
	for i := 0; i < 3; i++ {
		getBin()
	}
}
