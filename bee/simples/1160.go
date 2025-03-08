package main

import (
	"fmt"
	"math"
)

func get() string {
	var v1, v2, v3, v4 float64
	fmt.Scan(&v1, &v2, &v3, &v4)

	for i := 0; i < 100; i++ { // Correção no loop
		if v1 > v2 {
			return fmt.Sprintf("%v anos.", i) // Correção na formatação da string
		}

		if v3 != 0 {
			v1 = v1 * (1 + v3/100)
			v1 = math.Floor(v1)
		}

		if v4 != 0 {
			v2 = v2 * (1 + v4/100)
			v2 = math.Floor(v2)
		}
	}

	return "Mais de 1 seculo."
}

func main() {
	var voltas float64
	fmt.Scan(&voltas)

	for i := 0; i < int(voltas); i++ {
		fmt.Println(get())
	}
}
