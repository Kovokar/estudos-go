package main

import (
	"fmt"
	"io"
)

func seq(num int, cont int) {
	lista := []int{0}
	for i := 1; i <= num; i++ {
		for j := 0; j < i; j++ {
			lista = append(lista, i)
		}
	}
	if len(lista) == 1 {
		fmt.Printf("Caso %d: %d numero\n", cont, len(lista))
	} else {
		fmt.Printf("Caso %d: %d numeros\n", cont, len(lista))
	}

	for i := 0; i < len(lista); i++ {
		if i == len(lista)-1 {
			fmt.Println(lista[i]) // Último item sem espaço
		} else {
			fmt.Print(lista[i], " ") // Espaço entre os números
		}
	}
}

func main() {
	cont := 1
	for {
		var o int
		_, err := fmt.Scanf("%d", &o)

		if err == io.EOF {
			break
		}
		seq(o, cont)
		cont++
		fmt.Println()
	}
}
