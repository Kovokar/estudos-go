package main

import (
	"fmt"
)

var distancia float64
var consumo float64

func main() {

	fmt.Scan(&distancia)
	fmt.Scan(&consumo)

	fmt.Printf("%.3f km/l", distancia/consumo)

}
