package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// fmt.Println(subString("Hey This java is hot", "Java is a new paradigm"))

	exec()
}

func subString(v1 string, v2 string) int {
	max := 0
	for i := 0; i < len(v1); i++ {
		for k := 0; k < len(v2); k++ {
			if v1[i] == v2[k] && !verificaEspacoFinal(v1, v2) {
				ret := recur(v1, v2, i+1, k+1, 1)
				if ret > max {
					max = ret
				}
			}
		}
	}
	return max
}

func recur(v1, v2 string, idx1, idx2, cont int) int {
	// fmt.Println(idx1, idx2, cont)
	if idx1 >= len(v1) || idx2 >= len(v2) {
		return cont
	}

	if v1[idx1] == v2[idx2] {
		return recur(v1, v2, idx1+1, idx2+1, cont+1)
	}
	return cont
}

func exec() {
	reader := bufio.NewReader(os.Stdin)

	for {
		str1, _ := reader.ReadString('\n')
		str2, _ := reader.ReadString('\n')
		if str1 == "" || str2 == "" {
			break
		}
		fmt.Println(subString(str1, str2))

	}
}

func verificaEspacoFinal(palavra1, palavra2 string) bool {
	return strings.HasSuffix(palavra1, " ") && strings.HasSuffix(palavra2, " ")
}
