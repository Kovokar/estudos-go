package main

import "fmt"

func main() {
	anonimous := struct {
		nome  string
		idade int
	}{
		nome:  "Maria",
		idade: 45,
	}

	fmt.Println(anonimous)
}
