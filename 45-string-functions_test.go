package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ae = assert.Equal
var at = assert.True

func TestStringFunctions(t *testing.T) {
	at(t, strings.Contains("test", "es"))
	ae(t, 2, strings.Count("test", "t"))
	at(t, strings.HasPrefix("test", "te"))
	at(t, strings.HasSuffix("test", "st"))
	ae(t, 1, strings.Index("test", "e"))
	ae(t, "a-b", strings.Join([]string{"a", "b"}, "-"))
	ae(t, "aaaaa", strings.Repeat("a", 5))
	ae(t, "f00", strings.Replace("foo", "o", "0", -1))
	ae(t, "f0o", strings.Replace("foo", "o", "0", 1))
	ae(t, []string{"a", "b", "c", "d", "e"}, strings.Split("a-b-c-d-e", "-"))
	ae(t, "test", strings.ToLower("TEST"))
	ae(t, "TEST", strings.ToUpper("test"))
}
