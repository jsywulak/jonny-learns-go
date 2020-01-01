package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRandom(t *testing.T) {
	assert.Equal(t, 1, 1)

	fmt.Print(rand.Intn(100), ",")
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100), rand.Intn(100))

	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64() * 5)
	fmt.Println((rand.Float64() * 5) + 5)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	fmt.Println(r1.Intn(100))

}
