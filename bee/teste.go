package main

import (
	"fmt"
)

func main() {
	var aa, ab, ac string

	// Corrigido para usar %s %s
	fmt.Scanf("%s %s", &aa, &ab)
	ac += aa + " " + ab

	fmt.Println(aa)
	fmt.Println(ab)
	fmt.Println(ac)

}
