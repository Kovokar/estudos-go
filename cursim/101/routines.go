package main

import (
	"fmt"
	"time"
)

func numeros(done chan<- bool) {
	for i := 0; i < 26; i++ {
		fmt.Printf("%v ", i)
		time.Sleep(time.Millisecond * 100)
	}
	done <- true
}

func letrs(done chan<- bool) {
	for i := 'a'; i <= 'z'; i++ {
		fmt.Printf("%c ", i)
		time.Sleep(time.Millisecond * 100)
	}
	done <- true
}

func main() {
	cn := make(chan bool)
	cl := make(chan bool)

	go numeros(cn)
	go letrs(cl)

	<-cn
	<-cl
	// go sacanagem()
	// time.Sleep(time.Second * 5)

	defer fmt.Println("\nFim da execução")
}
