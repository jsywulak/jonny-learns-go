package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type response1 struct {
	Page   int
	Fruits []string
}
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruit"`
}

func TestJson(t *testing.T) {
	assert.Equal(t, 1, 1)
}
