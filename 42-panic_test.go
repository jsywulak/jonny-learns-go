package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPanic(t *testing.T) {
	assert.Equal(t, 2, 2)

	// panic("a problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
