package main

import "fmt"

func main() {
	nomes := map[string]string{
		"pedro":   "guilherme",
		"alessia": "vitória",
	}

	nomes["alice"] = "maria" // KKKKKKKKKK te fude alice, teu nome é maria

	fmt.Println(nomes)
	fmt.Println()

	delete(nomes, "alice")
	delete(nomes, "alessia")

	for k, v := range nomes {
		fmt.Println(k, ":", v)
	}
	println("Alone Again")

}
