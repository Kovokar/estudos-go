package main

import "fmt"

func main() {

	fmt.Println(isPalindrom("123456654321"))
}

func isPalindrom(str string) bool {
	ult := len(str) - 1
	cont := 0
	for i := 0; i < ult; i++ {
		cont++
		fmt.Println(cont)
		if str[i] != str[ult-i] {
			return false
		}
		if i == ult/2 {
			return true
		}
	}
	return true
}
