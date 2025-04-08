package main

import "fmt"

func main() {
	var arr = []int{4, 6, 8, 4, 52}
	printaArrat(arr)
	printaArray(arr)

}

func printaArrat(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func printaArray(arr []int) {
	for _, v := range arr {
		fmt.Println(v)
	}
}
