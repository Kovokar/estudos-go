package main

import "fmt"

func main() {
	fmt.Println(inverteArray([]int{1, 2, 3, 4, 5}))
}

func inverteArray(arr []int) (inverted []int) {
	for i := len(arr) - 1; i >= 0; i-- {
		inverted = append(inverted, arr[i])
	}
	return
}
