package main

import "fmt"

func main() {
	x := 10

	y := func(x int) {
		fmt.Println(x * x)
	}

	y(x)
	y(30)
}
