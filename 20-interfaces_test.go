package main

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

type geometry interface {
	area() float64
	perim() float64
}

type reckt struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r reckt) area() float64 {
	return r.height * r.width
}

func (r reckt) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println("dimensions:", g)
	fmt.Println("area:", g.area())
	fmt.Println("perimeter:", g.perim())
}

func TestInterfaces(t *testing.T) {
	assert.Equal(t, 1, 1)
	r := reckt{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	assert.Equal(t, float64(12), r.area())

	measure(c)
}
