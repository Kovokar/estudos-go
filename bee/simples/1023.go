package main

import "fmt"

var menor, voltas, consumoTotal, pessoasTotais int
var num1, num2, cont int
var slilce [][]int

// var subSlice []int

func main() {
	recebeValores()
	fmt.Println(slilce)
	// fmt.Println(consumoTotal)
	// fmt.Println(pessoasTotais)
	slilce = bubbleSort(slilce)
	fmt.Println(slilce)
	// mostrar()

}

func recebeValores() {
	fmt.Scan(&voltas)

	for i := 0; i < voltas; i++ {
		fmt.Scanf("%v %v", &num1, &num2)
		consumoTotal += num2
		pessoasTotais += num1
		slilce = append(slilce, []int{num1, num2 / num1})
	}
}

func bubbleSort(arr [][]int) [][]int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j][1] > arr[j+1][1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func mostrar() {
	for _, sub := range slilce {
		for i := 0; i < len(sub)-1; i++ {
			fmt.Printf("%v-%v ", sub[i], sub[i+1])
		}
	}
}
