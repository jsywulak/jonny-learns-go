package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func TestPointers(t *testing.T) {
	i := 1
	fmt.Println("initial:", i)
	assert.Equal(t, 1, i)

	zeroval(i)
	fmt.Println("zeroval", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer", &i)
}
