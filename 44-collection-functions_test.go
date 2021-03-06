package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollectionFunctions(t *testing.T) {
	var strs = []string{"peach", "apple", "pear", "plum"}
	assert.Equal(t, 2, Index(strs, "pear"))
	assert.False(t, Include(strs, "grape"))
	assert.True(t, Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	assert.False(t, All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	var expected = []string{"peach", "apple", "pear"}
	assert.Equal(t, expected, Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	expected = []string{"PEACH", "APPLE", "PEAR", "PLUM"}
	assert.Equal(t, expected, Map(strs, strings.ToUpper))
}

func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
