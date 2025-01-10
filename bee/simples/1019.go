package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	hora := N / 3600
	minuto := (N % 3600) / 60
	segundo := N % 60

	fmt.Printf("%d:%d:%d\n", hora, minuto, segundo)
}
