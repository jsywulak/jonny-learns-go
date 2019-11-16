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
	assert.Equal(t, 1, i)

	zeroval(i)
	assert.Equal(t, 1, i)

	zeroptr(&i)
	assert.Equal(t, 0, i)

	fmt.Println("pointer", &i)
	fmt.Printf("%T", &i)
	assert.NotEqual(t, 0, &i)
	assert.NotEqual(t, 1, &i)
}
