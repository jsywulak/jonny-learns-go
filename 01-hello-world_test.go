package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	helloworld := "Hello World!"
	assert.Equal(t, "Hello World!", helloworld)
}
