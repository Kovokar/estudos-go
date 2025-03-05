package main

import "fmt"

// Define uma função que aceita um callback (função como argumento)
func executarCallback(callback func(int, int) int) {
	resultado := callback(5, 3) // Chama a função callback com argumentos
	fmt.Println("Resultado do callback:", resultado)
}

func callback2(kall func(int) int) {

	a := kall(7)
	fmt.Println(a)
	fmt.Println(kall(9))
}

// Função que será passada como callback
func somar(a int, b int) int {
	return a + b
}

func elevaQuadrado(x int) int {
	return x * x
}

func main() {
	// Passa a função 'somar' como callback
	executarCallback(somar)
	callback2(elevaQuadrado)
}
