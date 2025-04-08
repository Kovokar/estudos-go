package main

import "fmt"

func main() {
	// Criando um slice com valores iniciais
	slice := []int{10, 20, 30, 40, 50}

	// Imprimindo o slice
	fmt.Println("Slice:", slice)

	// Adicionando um elemento ao slice
	slice = append(slice, 60)
	fmt.Println("Slice após append:", slice)

	// Modificando um valor no slice
	slice[2] = 100
	fmt.Println("Slice após modificação:", slice)

	// Sub-slice (slice de um slice)
	subSlice := slice[1:4] // Do índice 1 até 3 (4 não incluído)
	fmt.Println("Sub-slice:", subSlice)
}
