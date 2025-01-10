package main

import "fmt"

var a float64
var b float64
var c float64
var d float64
var e float64
var f float64

func main() {

	fmt.Scanf("%v", &a)
	fmt.Scanf("%v", &b)
	fmt.Scanf("%v", &c)
	fmt.Scanf("%v", &d)
	fmt.Scanf("%v", &e)
	fmt.Scanf("%v", &f)

	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", (b*c)+(e*f))
}
