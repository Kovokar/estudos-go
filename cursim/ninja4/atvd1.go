package main

import "fmt"

func main() {
	// var arr [5]int //{1, 2, 3, 4, 5}
	arr := [5]int{}

	fmt.Println(arr)

	for i, _ := range arr {
		arr[i] = i
	}

	fmt.Println(arr)
	fmt.Printf("Tipo: %T", arr)

}
