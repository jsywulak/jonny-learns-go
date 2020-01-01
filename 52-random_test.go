package main

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandom(t *testing.T) {
	assert.Equal(t, 1, 1)

	fmt.Print(rand.Intn(100), ",")
	fmt.Println(rand.Intn(100))
	fmt.Print(rand.Intn(100), ",", rand.Intn(100))
}
