package main

import "fmt"

func main() {
	// Array com tamanho fixo
	arr := [3]int{1, 2, 3}

	// Slice baseado no array
	slice := arr[:]

	// Modificando o slice
	slice[1] = 20

	fmt.Println("Array após modificação no slice:", arr)
	fmt.Println("Slice após modificação:", slice)
}
