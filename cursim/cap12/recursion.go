package main

import "fmt"

var gloval int = 1

func main() {
	fmt.Println(fatorial(6))
	fmt.Println(fibonacci(5))

}

func fatorial(i int) int {
	if i == 1 {
		return gloval
	} else {
		return i * fatorial(i-1)
	}
}

// Função recursiva para calcular Fibonacci
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
