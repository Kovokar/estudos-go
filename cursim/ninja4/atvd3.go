package main

import "fmt"

func main() {
	// var arr [5]int //{1, 2, 3, 4, 5}
	slc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(slc[:3])
	fmt.Println(slc[4:])
	fmt.Println(slc[3:7])
	fmt.Println(slc[4 : len(slc)-1])
	// fmt.Println()

}
