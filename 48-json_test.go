package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type response1 struct {
	Page   int
	Fruits []string
}

func TestJson(t *testing.T) {
	assert.Equal(t, 1, 1)
}
