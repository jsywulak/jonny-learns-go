package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type point struct {
	x, y int
}

func TestStringFormatting(t *testing.T) {
	assert.Equal(t, 1, 1)
	p := point{1, 2}
	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)
	fmt.Printf("%T\n", p)
	fmt.Printf("%d\n", 123)
}
