package main

import (
	"fmt"
)

func contaCarry(a, b int) int {
	carry := 0
	count := 0

	for a > 0 || b > 0 {
		soma := (a % 10) + (b % 10) + carry
		if soma >= 10 {
			carry = 1
			count++
		} else {
			carry = 0
		}
		a /= 10
		b /= 10
	}

	return count
}

func main() {
	for {
		var a, b int
		fmt.Scan(&a, &b)

		if a == 0 && b == 0 {
			break
		}

		carryCount := contaCarry(a, b)

		if carryCount == 0 {
			fmt.Println("No carry operation.")
		} else if carryCount == 1 {
			fmt.Println("1 carry operation.")
		} else {
			fmt.Printf("%d carry operations.\n", carryCount)
		}
	}
}
