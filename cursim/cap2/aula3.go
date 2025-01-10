package main

import "fmt"

func main() {
	x := 10
	y := "ainnn"
	z := true

	fmt.Printf("x: %v, %T\n", x, x)
	x = 20
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)
	fmt.Printf("z: %v, %T", z, z)

}
