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

	fmt.Println(r1.Intn(100), r1.Intn(100))

	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	r2_1 := r2.Intn(100)
	r2_2 := r2.Intn(100)

	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	r3_1 := r3.Intn(100)
	r3_2 := r3.Intn(100)
	assert.Equal(t, r2_1, r3_1)
	assert.Equal(t, r2_2, r3_2)

}
