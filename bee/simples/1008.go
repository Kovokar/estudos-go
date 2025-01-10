package main

import "fmt"

var x int
var y int
var z float64

func main() {

	fmt.Scanf("%v", &x)
	fmt.Scanf("%v", &y)
	fmt.Scanf("%v", &z)

	fmt.Printf("NUMBER = %v\n", x)
	fmt.Printf("SALARY = U$ %.2f\n", float64(y)*z)

}
