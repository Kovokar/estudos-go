package main

import "fmt"

// conversao de tipos

type hotdog int

var b hotdog = 20
var x int = 10

func main() {
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", x)

	x = int(b)
	fmt.Printf("%v\n", x)

}
