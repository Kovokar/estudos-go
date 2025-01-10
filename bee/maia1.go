// package main

// import (
// 	"fmt"
// )

// var nome string
// var nomes []string
// var total int
// var tamanhos []int

// func main() {

// 	exec()
// }

// func inputs(nome string, nomes []string, n int) []string {

// 	// n := 1
// 	// fmt.Scanf("%v", &n)
// 	for i := 0; i < n; i++ {
// 		fmt.Scanf("%v", &nome)
// 		nomes = append(nomes, nome)
// 	}
// 	return ordena_slice(nomes)
// }

// func ordena_slice(nomes []string) []string {

// 	tamanho_menos_um := len(nomes) - 1

// 	for i := 0; i < tamanho_menos_um; i++ {
// 		for j := 0; j < tamanho_menos_um-i; j++ {
// 			if len(nomes[j]) > len(nomes[j+1]) {
// 				temp := nomes[j]
// 				nomes[j] = nomes[j+1]
// 				nomes[j+1] = temp
// 			}
// 		}
// 	}

// 	return nomes
// }

// func b_contem_a(atual string, proxima string) bool {
// 	// tamanho_menos_um := len(nomes)
// 	pos_j := 0
// 	cont := 0
// 	voltas := 0

// 	// pertence := ""

// 	for i := 0; i < len(atual); i++ {
// 		letra_atual := string(atual[i])
// 		for j := pos_j; j < len(proxima); j++ {
// 			voltas++
// 			if letra_atual == string(proxima[j]) {
// 				pos_j = j + 1
// 				cont++
// 				break
// 			}
// 		}
// 	}

// 	// fmt.Println("LEN DO 1° -> ", len(atual))
// 	// fmt.Println("LEN DO 2° -> ", len(proxima))

// 	// fmt.Println("TOTAL DE VOLTAS: ", voltas)
// 	// fmt.Println("TOTAL DE IGUASI: ", cont)
// 	// fmt.Println("======================")

// 	return len(atual) == cont
// }

// func cultiva(nomes []string) int {

// 	// fmt.Println(nomes, "\n")

// 	atual := 0
// 	proxima := 1
// 	cont := 1

// 	for i := 0; i < len(nomes); i++ {

// 		if proxima == len(nomes) {
// 			return cont
// 		}
// 		if b_contem_a(nomes[atual], nomes[proxima]) {
// 			// fmt.Printf("A palavra %s está contida em %v\n", nomes[atual], nomes[proxima])
// 			cont++
// 			atual++
// 			proxima++
// 		} else {
// 			proxima++
// 		}
// 	}
// 	return cont
// }

// func exec() {
// 	var n_voltas int = 999
// 	for n_voltas != 0 {

// 		fmt.Scanf("%d", &n_voltas)
// 		nomes = nomes[:0]
// 		nomes = inputs(nome, nomes, n_voltas)
// 		if n_voltas != 0 {
// 			fmt.Println(cultiva(nomes))
// 			fmt.Println(nomes)
// 		}
// 	}
// }
