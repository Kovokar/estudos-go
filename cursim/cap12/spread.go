package main

import "fmt"

func main() {

	si := []int{1, 2, 3, 4, 6, 7, 8, 9}

	fmt.Println(soma(si...))

}

func soma(x ...int) int {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma

}
