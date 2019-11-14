package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValues(t *testing.T) {
	assert.Equal(t, "go"+"lang", "golang")

	assert.Equal(t, 1+1, 2)
	assert.Equal(t, 7.0/3.0, 2.3333333333333335)

	assert.Equal(t, true && false, false)
	assert.Equal(t, true || false, true)
	assert.Equal(t, !true, false)
}
