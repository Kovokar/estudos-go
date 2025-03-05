package main

import "fmt"

func callback2(kall func(int) int) {
	fmt.Println(kall(9))
}

// Função que será passada como callback
func elevaQuadrado(x int) int {
	return x * x
}

func main() {
	// Passa a função 'somar' como callback
	callback2(elevaQuadrado)
}
