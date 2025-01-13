package main

import (
	"fmt"
)

func main() {
	for x := 33; x <= 122; x++ {
		fmt.Printf("%d - %v ", x, string(x))
		fmt.Printf("%#x - %v ", x, string(x))
		fmt.Printf("%#U - %v\n", x, string(x))
	}
}
