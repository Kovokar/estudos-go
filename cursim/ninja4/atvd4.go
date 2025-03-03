package main

import "fmt"

func main() {
	// var arr [5]int //{1, 2, 3, 4, 5}
	slc()
	fmt.Println("COM MAKE:")
	makef()

}

func slc() {
	slc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slc)
	slc = append(slc, 10)
	fmt.Println(slc)
	slc2 := []int{11, 12, 13}
	slc = append(slc, slc2...)
	fmt.Println(slc)
}

func makef() {
	slc := make([]int, 0, 20)
	slc2 := []int{11, 12, 13}
	slc = append(slc, 2)
	slc = append(slc, slc2...)
	fmt.Println(slc)

}
