package main

import "fmt"

var numbs = []int{1, 2, 3, 4, 5}

func main() {
	fmt.Println(somaImpares(somaVals, numbs...))
}

func somaVals(i ...int) int {
	init := 0
	for _, v := range i {
		init += v
	}
	return init
}

func somaImpares(fun func(i ...int) int, list ...int) int {
	pares := make([]int, 10, 10)
	for _, v := range list {
		if v%2 == 1 {
			pares = append(pares, v)
		}
	}
	return fun(pares...)
}
