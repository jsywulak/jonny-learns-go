package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPanic(t *testing.T) {
	assert.Equal(t, 2, 2)

	// panic("a problem")
}
