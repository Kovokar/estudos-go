package main

import (
	"fmt"
)

func main() {
	var digito byte
	var valor string

	for {
		fmt.Scanf("%c", &digito)
		fmt.Scanf("%s", &valor)

		if digito == '0' {
			break
		}

		n := len(valor)
		for i := 0; i < n; i++ {
			if valor[i] == digito {
				valor = valor[:i] + valor[i+1:]
				n--
				i--
			}
		}

		for i := 0; i < n-1; i++ {
			if valor[i] == '0' {
				valor = valor[:i] + valor[i+1:]
				n--
				i--
			} else {
				break
			}
		}

		if n == 0 {
			valor = "0"
		}

		fmt.Println(valor)
	}
}
