package main

import (
	"fmt"
	"time"
)

func numeros() {
	for i := 0; i < 26; i++ {
		fmt.Printf("%v ", i)
		time.Sleep(time.Millisecond * 100)
	}
}

func letrs() {
	for i := 'a'; i <= 'z'; i++ {
		fmt.Printf("%c ", i)
		time.Sleep(time.Millisecond * 100)
	}
}

func sacanagem() {
	time.Sleep(time.Millisecond * 100)

	fmt.Println()
}

func main() {
	go numeros()
	go letrs()
	// go sacanagem()
	// time.Sleep(time.Second * 5)

	defer fmt.Println("\nFim da execução")
}
