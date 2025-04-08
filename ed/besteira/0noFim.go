package main

import "fmt"

func main() {
	fmt.Println(moveZeros([]int{0, 1, 0, 3, 12}))
}

func moveZeros(arr []int) []int {
	j := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 && arr[j] == 0 {
			arr[j], arr[i] = arr[i], arr[j]
		}
		if arr[j] != 0 {
			j++
		}
	}
	return arr
}
