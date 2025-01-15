package main

import (
	"fmt"
)

func main() {
	exec()
}

func soma3(palavra string) string {
	temp := ""
	for _, i := range palavra {

		if isLower(int(i)) {
			temp += (string(i + 3))
		} else {
			temp += (string(i))
		}
	}
	return temp
}

func isLower(i int) bool {
	return (i >= 97 && i <= 122) || (i >= 65 && i <= 90)
}

func inverte(palavra string) string {
	temp := ""
	for i := len(palavra) - 1; i != -1; i-- {
		temp += (string(palavra[i]))
	}
	return temp
}

func soma1(palavra string) string {
	temp := ""
	for i := 0; i < len(palavra); i++ {
		if i >= int(len(palavra)/2) {
			temp += (string(palavra[i] - 1))
		} else {
			temp += (string(palavra[i]))
		}
	}
	return temp
}

func exec() {
	num := 0
	fmt.Scan(&num)

	for i := 0; i < num; i++ {
		var palavra string
		fmt.Scan(&palavra)
		palavra = soma3(palavra)
		palavra = inverte(palavra)
		fmt.Println(soma1(palavra))
	}
}
