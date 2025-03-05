package main

import "fmt"

func main() {
	m1 := multiplicador()
	fmt.Println(m1(2))
	fmt.Println(m1(3))
	fmt.Println(m1(4))

}

func multiplicador() func(i int) int {
	tot := 1

	return func(i int) int {
		return (tot * i)
	}
}
