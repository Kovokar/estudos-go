package main

import "fmt"

func main() {
	var tempo, velocidade float64
	fmt.Scan(&tempo)
	fmt.Scan(&velocidade)
	fmt.Printf("%.3f\n", tempo*velocidade/12)
}
