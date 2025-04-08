package main

import "fmt"

func main() {
	// var a [5]int
	// var b [5]int
	// var c [5]int
	// var d [5]int
	// var e [5]int
	// var f [5]int
	// var g [5]int
	// var h [5]int
	// var i [5]int
	// var j [5]int
	// fmt.Println(a, b, c, d, e, f, g, h, i, j)
	var b = []int{1, 2, 3}
	b = []int{4, 5, 6}
	var c = [...]int{7, 8, 9}
	var d = [...]int{10, 11, 12}
	var e = [...]int{13, 14, 15}
	var f = [...]int{16, 17, 18}
	var g = [...]int{19, 20, 21}
	var h = [...]int{22, 23, 24}
	var i = [...]int{25, 26}

	// fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
}
