package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

const s string = "constant"

func TestConstants(t *testing.T) {
	assert.Equal(t, "constant", s)

	const n = 500000000
	const d = 3e20 / n
	assert.Equal(t, 6e+11, d)
	assert.Equal(t, int64(600000000000), int64(d))
	assert.Equal(t, -0.28470407323754404, math.Sin(n))
}
