package main

import "fmt"

func main() {
	fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println(4)
	fmt.Println(5)
	fmt.Println(6)

}
