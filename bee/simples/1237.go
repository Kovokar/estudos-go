package main

import (
	"fmt"
)

func main() {
	subString("abcdef", "cdofhij")
}

func subString(v1 string, v2 string) int {
	for i := 0; i < len(v1); i++ {
		for k := 0; k < len(v2); k++ {
			if v1[i] == v2[k] {
				fmt.Println("chamou")
				ret := recur(v1, v2, i, k, 0)
				fmt.Println("wqeqwe", ret)
			}
		}
	}
	return 0
}

func recur(v1, v2 string, idx1, idx2, cont int) int {
	// fmt.Println(idx1, idx2)
	if v1[idx1] == v2[idx2] {
		fmt.Println(string(v1[idx1]), string(v2[idx2]))
		recur(v1, v2, idx1+1, idx2+1, cont+1)
	}
	if cont == 100 {
		return 99999
	}
	fmt.Println("dfwq", cont)
	return cont
}
