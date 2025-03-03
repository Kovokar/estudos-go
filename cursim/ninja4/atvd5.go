package main

import "fmt"

func main() {
	// var arr [5]int //{1, 2, 3, 4, 5}
	slc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slc)

}

func slc() {
	slc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slc)
	slc2 := append(slc[:3], slc[7:]...)
	fmt.Println(slc2)
}
