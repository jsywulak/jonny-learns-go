package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaps(t *testing.T) {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	assert.Equal(t, len(map[string]int{"k1": 7, "k2": 13}), len(m))

	assert.Equal(t, 7, m["k1"])
	assert.Equal(t, 2, len(m))
	delete(m, "k2")
	assert.Equal(t, 1, len(m))

	_, prs := m["k2"]
	fmt.Println("prs:", prs)
	assert.False(t, prs, "no error should be returned")

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
