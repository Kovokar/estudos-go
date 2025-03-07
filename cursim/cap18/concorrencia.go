package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	wg.Add(2)
	go f1()
	go f2()

	fmt.Println(runtime.NumGoroutine())
	wg.Wait()
}

func f1() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
		time.Sleep(20)
	}
	wg.Done()
}

func f2() {
	for i := 99; i >= 0; i-- {
		fmt.Println(i, "xxxx")
		time.Sleep(20)
	}
	wg.Done()

}
