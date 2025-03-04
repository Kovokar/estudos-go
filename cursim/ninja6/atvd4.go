package main

import "fmt"

func main() {
	anon := struct {
		pri map[string]string
		seg []int
	}{
		pri: map[string]string{
			"alfredo": "5551234",
			"joana":   "9996674",
		},
		seg: []int{1, 2, 3, 4},
	}

	fmt.Println(anon)
}
