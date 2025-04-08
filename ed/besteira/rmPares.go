package main

import "fmt"

func main() {
	fmt.Println(rmPares([]int{1, 2, 3, 4, 5}))
}

func rmPares(nums []int) (odds []int) {
	for _, v := range nums {
		if v%2 != 0 {
			odds = append(odds, v)
		}
	}
	return
}
