package main

import "fmt"

// conversao de tipos

type sub int

var x sub

func main() {

	fmt.Printf("x = %v\n", x)
	fmt.Printf("x = %T\n", x)
	x = 42
	fmt.Printf("x = %v\n", x)
}
