package main

import "fmt"

func main() {

	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("4")
	fmt.Println("5")
	fmt.Println("6")
	fmt.Println("7")
	fmt.Println("8")
	fmt.Println("9")
	fmt.Println("10")

}
