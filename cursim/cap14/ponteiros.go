package main

import "fmt"

func main() {
	x := 10

	y := &x

	fmt.Println(*y)
	*y++
	fmt.Println(*y)

}
