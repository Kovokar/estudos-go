package main

import "fmt"

func main() {
	fmt.Println(length("abaco"))
	fmt.Println(soma(1, 2, 3, 4, 5, 6))
}

func length(str string) int {
	cont := 0
	for range str {
		cont++
	}
	return cont
}

func soma(x ...int) int {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma

}
