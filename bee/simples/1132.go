package main

import "fmt"

func main() {
	var v1, v2 int
	fmt.Scan(&v1)
	fmt.Scan(&v2)

	if v1 > v2 {
		v1, v2 = v2, v1
	}

	v := 0
	for i := v1; i <= v2; i++ {
		if !(i%13 == 0) {
			v = v + i
		}
	}
	fmt.Println(v)
}
