package main

import "fmt"

func main() {
	// var arr [5]int //{1, 2, 3, 4, 5}
	slc := []int{9, 8, 7, 6, 5}

	fmt.Println(slc)

	for i, _ := range slc {
		slc[i] = i
	}

	fmt.Println(slc)
	fmt.Printf("Tipo: %T", slc)

}
