package main

import "fmt"

func main() {
	x := 2
	func(x int) {
		fmt.Println(x)
	}(x)
}
