package main

import "fmt"

func main() {
	x := 10

	func(x int) {
		fmt.Println(x * x)
	}(x)
}
