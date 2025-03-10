package main

import (
	"fmt"
)

func menore() error {
	menor := 9999999999999999999999999999.9
	var voltas int
	_, err := fmt.Scanf("%d", &voltas)
	if err != nil {
		return err
	}

	for i := 0; i < voltas; i++ {
		var num float64
		fmt.Scanf("%f", &num)
		if num < menor {
			menor = num
		}
	}

	fmt.Printf("%.2f\n", menor)

	return nil
}

func main() {
	for {
		if err := menore(); err != nil {
			break
		}
	}
}
