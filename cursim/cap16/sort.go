package main

type Carro struct {
	nome     string
	potencia int
	consumo  int
}

type Interface interface {
	// Len is the number of elements in the collection.
	Len() int

	Less(i, j int) bool

	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

func main() {

}
