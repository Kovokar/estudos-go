package main

import "fmt"

func main() {
	var dias int
	fmt.Scan(&dias)

	anos := dias / 365
	dias = dias % 365
	meses := dias / 30
	dias = dias % 30

	fmt.Printf("%d ano(s)\n", anos)
	fmt.Printf("%d mes(es)\n", meses)
	fmt.Printf("%d dia(s)\n", dias)
}
