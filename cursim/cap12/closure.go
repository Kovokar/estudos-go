package main

import "fmt"

// função somador recebe como paramentro nada e retorna
// uma função que por sua vez recebe um int e retora um int
func somador() func(r int) float64 {

	soma := 0

	return func(i int) float64 {
		soma = soma + i
		return float64(soma)
	}

}

func main() {

	a1 := somador()
	fmt.Println(a1(2))
	fmt.Println(a1(2))
	fmt.Println(a1(2))
	fmt.Println("------")
	a2 := somador()
	fmt.Println(a2(2))
	fmt.Println(a2(5))
	fmt.Println(a2(3))
}
