package main

import (
	"fmt"
)

func get() (int, int, int) {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	return a, b, c
}

func cases(a, b, c int) {
	// Desceu e subiu ou igual
	if a > b && b <= c {
		fmt.Println(":)")
		return
	}
	// Subiu e desceu ou igual
	if a < b && b >= c {
		fmt.Println(":(")
		return
	}
	// Subiu e depois subiu menos
	if a < b && b < c && (b-a > c-b) {
		fmt.Println(":(")
		return
	}
	// Subiu e depois subiu igual ou mais
	if a < b && b < c && (b-a <= c-b) {
		fmt.Println(":)")
		return
	}
	// Desceu e depois desceu menos
	if a > b && b > c && (a-b > b-c) {
		fmt.Println(":)")
		return
	}
	// Desceu e depois desceu igual ou mais
	if a > b && b > c && (a-b <= b-c) {
		fmt.Println(":(")
		return
	}
	// Temperatura constante do 1º para o 2º dia
	if a == b {
		if b < c {
			fmt.Println(":)")
		} else {
			fmt.Println(":(")
		}
	}
}

func main() {
	a, b, c := get()
	cases(a, b, c)
}
