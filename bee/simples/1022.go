package main

import "fmt"

var N1, N2, D1, D2 int
var R1, R2 int
var operando string

func main() {
	exec()

}

func retornaOperando() {

	fmt.Scanf("%v / %v %s %v / %v", &N1, &D1, &operando, &N2, &D2)
	switch operando {
	case "+":
		R1 = (N1*D2 + N2*D1)
		R2 = (D1 * D2)
	case "-":
		R1 = (N1*D2 - N2*D1)
		R2 = (D1 * D2)
	case "*":
		R1 = (N1 * N2)
		R2 = (D1 * D2)
	case "/":
		R1 = (N1 * D2)
		R2 = (N2 * D1)
	default:
		fmt.Println("Operação inválida")
	}

}

func mdc(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func simplificar(r1, r2 int) (int, int) {
	divisor := mdc(r1, r2)
	return r1 / divisor, r2 / divisor
}

func exec() {
	var voltas int
	fmt.Scan(&voltas)
	for i := 0; i < voltas; i++ {
		retornaOperando()
		fmt.Printf("%v/%v", R1, R2)
		R1, R2 = simplificar(R1, R2)
		if R2 < 0 {
			R2 = R2 * -1
			R1 = R1 * -1
		}
		fmt.Printf(" = %v/%v\n", R1, R2)
	}
}
