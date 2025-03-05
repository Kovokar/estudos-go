package main

import "fmt"

func main() {
	x := func(e int) int {
		return e
	}(4)

	fmt.Println(x)
}
