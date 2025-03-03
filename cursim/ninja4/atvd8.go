package main

import "fmt"

func main() {
	nomes := map[string]string{
		"pedro":   "guilherme",
		"alessia": "vitória",
	}

	fmt.Println(nomes)
	fmt.Println(nomes["pedro"])

}
