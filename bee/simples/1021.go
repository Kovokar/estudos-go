package main

import "fmt"

func main() {
	var N float64
	fmt.Scan(&N)
	notaValores := []float64{100.00, 50.00, 20.00, 10.00, 5.00, 2.00}
	moedaValores := []float64{1.00, 0.50, 0.25, 0.10, 0.05, 0.01}

	intValor := int(N * 100)

	fmt.Println("NOTAS:")
	for _, valor := range notaValores {
		quantidade := intValor / int(valor*100)
		intValor %= int(valor * 100)
		fmt.Printf("%d nota(s) de R$ %.2f\n", quantidade, valor)
	}

	fmt.Println("MOEDAS:")
	for _, valor := range moedaValores {
		quantidade := intValor / int(valor*100)
		intValor %= int(valor * 100)
		fmt.Printf("%d moeda(s) de R$ %.2f\n", quantidade, valor)
	}
}
