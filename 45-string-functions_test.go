package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var p = fmt.Println

func TestStringFunctions(t *testing.T) {
	assert.True(t, strings.Contains("test", "es"))
	assert.Equal(t, 2, strings.Count("test", "t"))
	assert.True(t, strings.HasPrefix("test", "te"))
	assert.True(t, strings.HasSuffix("test", "st"))
	assert.Equal(t, 1, strings.Index("test", "e"))
	assert.Equal(t, "a-b", strings.Join([]string{"a", "b"}, "-"))
	assert.Equal(t, "aaaaa", strings.Repeat("a", 5))
	assert.Equal(t, "f00", strings.Replace("foo", "o", "0", -1))
	assert.Equal(t, "f0o", strings.Replace("foo", "o", "0", 1))
	assert.Equal(t, []string{"a", "b", "c", "d", "e"}, strings.Split("a-b-c-d-e", "-"))
	assert.Equal(t, "test", strings.ToLower("TEST"))
	assert.Equal(t, "TEST", strings.ToUpper("test"))
}
