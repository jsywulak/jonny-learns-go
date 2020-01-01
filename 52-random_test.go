package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRandom(t *testing.T) {

	for i := 0; i < 1000; i++ {
		randInt := rand.Intn(100)
		assert.True(t, 0 <= randInt && randInt < 100)
		randFloat1 := rand.Float64()
		assert.True(t, 0.0 <= randFloat1 && randFloat1 < 1.0)
		randFloat2 := (rand.Float64() * 5)
		assert.True(t, 0.0 <= randFloat2 && randFloat2 < 5.0)
		randFloat3 := ((rand.Float64() * 5) + 5)
		assert.True(t, 5.0 <= randFloat3 && randFloat3 < 10.0)
	}
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
