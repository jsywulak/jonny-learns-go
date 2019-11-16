package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func TestMethods(t *testing.T) {
	r := rect{width: 10, height: 5}

	assert.Equal(t, 50, r.area())
	assert.Equal(t, 30, r.perim())

	rp := &r

	assert.Equal(t, 50, rp.area())
	assert.Equal(t, 30, rp.perim())
}
