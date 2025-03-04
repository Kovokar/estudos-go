package main

import "fmt"

type Cliente struct {
	nome      string
	sobrenome string
	fumante   bool
}

func main() {

	cliente1 := Cliente{
		nome:      "Jõao",
		sobrenome: "fenelon",
		fumante:   false,
	}

	cliente2 := Cliente{"saath", "Hawury", true}

	fmt.Println(cliente1)
	fmt.Println(cliente2)

}
