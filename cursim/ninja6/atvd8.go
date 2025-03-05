package main

import "fmt"

func main() {
	soma := retornaFunc(20)
	fmt.Println(soma())
}

func retornaFunc(x int) (f func() int) {
	soma := x + 10

	somaFunc := func() int {
		return soma
	}
	return somaFunc
}
