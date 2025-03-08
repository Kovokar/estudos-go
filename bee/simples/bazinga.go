package main

import "fmt"

func getVals(v1, v2 string) string {
	if v1 == v2 {
		return "De novo!"
	} else if (v1 == "tesoura" && (v2 == "papel" || v2 == "lagarto")) ||
		(v1 == "papel" && (v2 == "pedra" || v2 == "Spock")) ||
		(v1 == "pedra" && (v2 == "lagarto" || v2 == "tesoura")) ||
		(v1 == "lagarto" && (v2 == "Spock" || v2 == "papel")) ||
		(v1 == "Spock" && (v2 == "tesoura" || v2 == "pedra")) {
		return "Bazinga!"
	} else {
		return "Raj trapaceou!"
	}
}

func main() {
	var o int
	fmt.Scan(&o)

	for i := 0; i < o; i++ {
		var v1, v2 string
		fmt.Scan(&v1, &v2)
		fmt.Printf("Caso #%d: %s\n", i+1, getVals(v1, v2))
	}
}
