package main

import "fmt"

var pi float64 = 3.14159
var raio float64
var area float64

func main() {

	fmt.Scanf("%v", &raio)
	area = (raio * raio) * pi

	fmt.Printf("A=%.4f\n", area)

}
