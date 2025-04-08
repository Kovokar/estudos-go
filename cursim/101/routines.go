package main

import (
	"fmt"
	"time"
)

func numeros(n chan<- int) {
	for i := 0; i < 26; i++ {
		// fmt.Printf("%v ", i)
		n <- i
		fmt.Printf("escrito do chan %d\n", i)
		time.Sleep(time.Millisecond * 100)
	}
	close(n)
	// done <- true
}

// func letrs(done chan<- bool) {
// 	for i := 'a'; i <= 'z'; i++ {
// 		fmt.Printf("%c ", i)
// 		time.Sleep(time.Millisecond * 100)
// 	}
// 	done <- true
// }

func main() {
	cn := make(chan int)
	go numeros(cn)

	for v := range cn {
		fmt.Printf("lido do chan %d\n ", v)
		time.Sleep(time.Millisecond * 90)

	}
	// go sacanagem()
	// time.Sleep(time.Second * 5)

	defer fmt.Println("\nFim da execução")
}
