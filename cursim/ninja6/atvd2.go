package main

import "fmt"

var list = []int{1, 2, 3, 4}

func main() {
	fmt.Println(spread(1, 2, 3, 4, 5))
	fmt.Println(sumList(list))
}

func spread(i ...int) int {
	tot := 0
	for _, v := range i {
		tot += v
	}
	return tot
}

func sumList(list []int) int {
	sum := 0
	for _, v := range list {
		sum += v
	}
	return sum
}
