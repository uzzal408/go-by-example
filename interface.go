package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}
type circle struct {
	redious float64
}

func (c *circle) area() float64 {
	return math.Pi * c.redious * c.redious
}

func (c *circle) perim() float64 {
	return 2 * math.Pi * c.redious
}

func mesure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
