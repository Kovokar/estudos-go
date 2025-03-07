package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	senha := "12345678"

	sb, err := bcrypt.GenerateFromPassword([]byte(senha), 10)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(sb))

}
