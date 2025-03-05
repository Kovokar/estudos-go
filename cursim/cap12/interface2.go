package main

import (
	"fmt"
	"math"
)

type Triangulo struct {
	tipo   string
	base   float64
	altura float64
}

type Quadrado struct {
	tipo string
	lado float64
}

type Circulo struct {
	tipo string
	raio float64
}

type Principal interface {
	area()
}

func (t Triangulo) area() {
	fmt.Println("Oi, sou do tipo ", t.tipo)
	fmt.Println("Minha área é ", (t.base*t.altura)/2)
}

func (q Quadrado) area() {
	fmt.Println("Oi, sou do tipo ", q.tipo)
	fmt.Println("Minha área é ", q.lado*q.lado)
}

func (c Circulo) area() {
	fmt.Println("Oi, sou do tipo ", c.tipo)
	fmt.Println("Minha área é ", c.raio*c.raio*math.Pi)

}

func mostrar(p Principal) {
	p.area()
}

func main() {
	triangulo1 := Triangulo{"triangulo", 4.0, 7}
	quadrado1 := Quadrado{"Quadrado", 7}
	circulo1 := Circulo{"Circulo", 5}

	mostrar(triangulo1)
	mostrar(quadrado1)
	mostrar(circulo1)

}
