package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	cedulas := []int{100, 50, 20, 10, 5, 2, 1}
	fmt.Println(N)
	for _, c := range cedulas {
		quantidade := N / c
		N = N % c
		fmt.Printf("%d nota(s) de R$ %d,00\n", quantidade, c)
	}
}
