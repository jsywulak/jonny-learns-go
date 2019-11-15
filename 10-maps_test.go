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

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
