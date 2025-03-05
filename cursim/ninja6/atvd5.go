package main

import (
	"fmt"
	"math"
)

type Quadrado struct {
	forma string
	lado  float64
}

type Circulo struct {
	forma string
	raio  float64
}

type Figura interface {
	area()
}

func (q Quadrado) area() {
	fmt.Println(q.forma, "\n", q.lado*q.lado)
}

func (c Circulo) area() {
	fmt.Println(c.forma, "\n", c.raio*c.raio*math.Pi)
}

func info(f Figura) {
	f.area()
}

func main() {
	quadrado1 := Quadrado{"Quadrado", 7}
	circulo1 := Circulo{"Circulo", 5}

	info(quadrado1)
	info(circulo1)
}
