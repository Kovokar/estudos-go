package main

import "fmt"

func main() {
	fmt.Println(inverteArray([]int{1, 2, 3, 4, 5}))
	fmt.Println(inverteArray2([]int{1, 2, 3, 4, 5}))
}

func inverteArray(arr []int) (inverted []int) {
	for i := len(arr) - 1; i >= 0; i-- {
		inverted = append(inverted, arr[i])
	}
	return
}

func inverteArray2(arr []int) []int {
	start := 0
	end := len(arr) - 1

	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
	return arr
}
